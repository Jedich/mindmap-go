<template>
	<svg class="canvas"></svg>
</template>

<script>
import * as d3 from "d3";
import { mapActions, mapGetters } from "vuex";

const imgWidth = 180

const props = {
	data: {
		type: Object,
		required: false
	},
}

export default {
	props,
	watch: {
		data(newValue) {
			this.launch();
		}
	},
	computed: {
		...mapGetters("select", {
			getCurrentNode: "getCurrentNode",
		}),
		...mapGetters("maps", {
			getCurrentMap: "getCurrentMap",
		}),
	},
	mounted() {
		this.launch();
	},
	methods: {
		...mapActions("select", {
			select: "select",
			deselect: "deselect"
		}),
		launch() {
			const width = 500;
			const data = this.data;
			var diagonal = d3.linkHorizontal().x(d => d.y).y(d => d.x)
			const root = d3.hierarchy(data);
			var dy = width
			var dx = 80;

			root.x0 = 0;
			root.y0 = dy / 2;
			root.descendants().forEach((d, i) => {
				d.id = i;
				d._children = d.children;
				//if (d.depth) d.children = null;
			});

			const svg = d3.select("svg")
				.attr('width', "100%")
				.attr('height', "100%")
				.style("font", "10px sans-serif")
				.style("user-select", "none");

			const g = svg.append("g")
			//.attr("transform", `translate(${width / 2},${margin.top})`);

			const gLink = g.append("g")
				.attr("fill", "none")
				.attr("stroke", "#555")
				.attr("stroke-opacity", 0.4)
				.attr("stroke-width", 1.5);

			const gNode = g.append("g")
				.attr("cursor", "pointer")
				.attr("pointer-events", "all");

			this.internaldata = {
				svg,
				g,
				gLink,
				gNode,
				root,
				dx,
				dy,
				diagonal
			}

			const zoomBehaviours = d3.zoom()
				.scaleExtent([0.05, 3])
				.on('zoom', (event) => g.attr('transform', event.transform))
				.filter((event) => (event.button === 1) || event.type === 'wheel');

			svg.call(zoomBehaviours);

			setTimeout(() => zoomBehaviours.translateTo(svg, 0, 0), 100);

			// Returns path data for a rectangle with rounded right corners.
			// The top-left corner is ⟨x,y⟩.
			function rightRoundedRect(x, y, width, height, radius) {
				return "M" + x + "," + y
					+ "h" + (width - radius)
					+ "a" + radius + "," + radius + " 0 0 1 " + radius + "," + radius
					+ "v" + (height - 2 * radius)
					+ "a" + radius + "," + radius + " 0 0 1 " + -radius + "," + radius
					+ "h" + (radius - width)
					+ "z";
			}

			function zoomToFit(paddingPercent) {
				const bounds = g.node().getBBox();
				const parent = svg.node().parentElement;
				const fullWidth = parent.clientWidth;
				const fullHeight = parent.clientHeight;

				const width = bounds.width;
				const height = bounds.height;

				const midX = bounds.x + (width / 2);
				const midY = bounds.y + (height / 2);

				if (width == 0 || height == 0) return; // nothing to fit

				const scale = (paddingPercent || 0.75) / Math.max(width / fullWidth, height / fullHeight);
				const translate = [fullWidth / 2 - scale * midX, fullHeight / 2 - scale * midY];

				const transform = d3.zoomIdentity
					.translate(translate[0], translate[1])
					.scale(scale);

				svg
					.transition()
					.duration(500)
					.call(zoomBehaviours.transform, transform);
			}

			this.update(root);

			return svg.node();
		},
		a() {
			this.internaldata.root.x0 = 0;
			this.internaldata.root.y0 = this.internaldata.dy / 2;
			this.internaldata.root.descendants().forEach((d, i) => {
				d.id = i;
				d._children = d.children;
				//if (d.depth) d.children = null;
			});
		},
		updateFromSky(data) {
			var temp = this.getCurrentNode.s.data()[0]
			//this.internaldata.root = d3.hierarchy(this.data);
			this.insert(temp, data)
			this.a()
			var newNode = this.update(temp)
			this.clickNode(null, newNode.data()[0])
		},
		updateSelected() {
			this.nodeText(this.getCurrentNode.s.select('text'));
			this.nodeImg(this.getCurrentNode.s.select('image'));
			this.a()
			this.update(this.getCurrentNode.s)
			this.wrapText(this.getCurrentNode.s);
		},
		nodeText(selection) {
			selection
				.attr("x", 0)
				.attr("y", d => { d.data.wrappedText = this.wrapRecu(d.data.name, 13); return -15 - (d.data.wrappedText.length - 1) * 10 })
				.attr("dy", "0em")
				.attr("font-size", 20)
				.text(d => d.data.wrappedText.join("/"));
		},
		nodeImg(selection) {
			selection
				.attr('xlink:href', d => {console.log(d.data.file ? d.data.file.filename : ""); return d.data.file ? `/api/img/${d.data.file.filename}` : ""})
				.attr('x', d =>  d.data.file ? -imgWidth/2 : null)
				.attr('y', d =>  d.data.file ? -(d.data.file.height*imgWidth/d.data.file.width)/2 -20: null)
				.attr('width', d => d.data.file ? imgWidth : null)
				.attr('height', d => d.data.file ? d.data.file.height*imgWidth/d.data.file.width: null)
				.attr("style", "overflow: hidden;");
		},
		insert(par, newNodeData) {
			console.log(par)
			console.log(this.getCurrentMap)
			newNodeData.created = true;
			let newNode = d3.hierarchy(newNodeData);
			newNode.depth = par.depth + 1;
			newNode.parent = par;
			newNode.created = true;
			if (!par.children)
				par.children = [];
			par.children.push(newNode);
			par.data.children.push(newNodeData)
			par._children = par.children;
			this.getCurrentNode.s.select('rect').style("fill", "#fff")
		},
		clickNode(event, d) {
			var sel = d3.selectAll('g.node')
				.filter(function (event, d2) {
					return d.id == this.id;
				})
			var thisNode = {
				id: sel.data()[0].id,
				s: sel,
				data: sel.data()[0].data
			}
			if (this.getCurrentNode === null) {
				this.select(thisNode);
			} else if (this.getCurrentNode.id === thisNode.id) {
				this.deselect();
			} else {
				this.deselect();
				this.select(thisNode);
			}
		},
		wrapText(nodeEnter) {
			this.wrap(nodeEnter.selectAll('text'), 5);
		},
		wrapRecu(text, limit, resultArr = []) {
			if (text.length > limit) {
				// find the last space within limit
				var line = text.slice(0, limit)
				var len = line.length
				var arr = [' ', '\n']
				arr.forEach(element => {
					var edge = line.lastIndexOf(element);
					if (edge > 0) {
						line = text.slice(0, edge);
						len = edge + 1
					}
				});
				var remainder = text.slice(len);
				resultArr.push(line)
				return this.wrapRecu(remainder, limit, resultArr);
			} else {
				resultArr.push(text)
			}
			if (resultArr.length === 0) {
				return [text]
			}
			return resultArr;
		},

		wrap(text, width) {
			text.each(function () {
				var text = d3.select(this),
					words = text.text().split("/"),
					lineNumber = 0,
					lineHeight = 1, // ems
					x = text.attr("x"),
					y = text.attr("y"),
					dy = parseFloat(text.attr("dy")),
					t = text.attr("y", 500),
					tspan = text.text(null).append("tspan").attr("text-anchor", "middle").attr("x", x).attr("y", y).attr("dy", dy + "em");
					d = text.data()[0];
					h = (d.data.file ? d.data.file.height*imgWidth/d.data.file.width : 0)
				words.forEach(word => {
					tspan = text.append("tspan").attr("x", x).attr("text-anchor", "middle").attr("y", y).attr("dy", ++lineNumber * lineHeight + dy + h/38 + "em").text(word);
				});
				// find corresponding rect and reszie
				var d = text.data()[0];
				var h = 50 + ((lineNumber - 1) * 19) + (d.data.file ? d.data.file.height*imgWidth/d.data.file.width : 0)
				var img = d3.select(this.parentNode).select('image')
				img.attr('y', -h / 2 + 5)
				d3.select(this.parentNode.children[0]).attr('height', h).attr('y', -h / 2);

			});
		},
		update(source) {
			//console.log("invoked")
			//console.log(this.internaldata.root)
			const duration = d3.event && d3.event.altKey ? 2500 : 250;
			var root = this.internaldata.root
			var svg = this.internaldata.svg
			var g = this.internaldata.g
			var gLink = this.internaldata.gLink
			var gNode = this.internaldata.gNode
			var dx = this.internaldata.dx
			var dy = this.internaldata.dy
			var diagonal = this.internaldata.diagonal

			const nodes = root.descendants().reverse();
			const links = root.links();


			var tree = d3.tree().nodeSize([dx, dy])
			// Compute the new tree layout.
			tree(root);

			const transition = svg.transition()
				.duration(duration)
				.tween("resize", window.ResizeObserver ? null : () => () => svg.dispatch("toggle"));

			// Update the nodes…
			const node = gNode.selectAll("g")
				.data(nodes, d => d.id)

			// Enter any new nodes at the parent's previous position.
			const nodeEnter = node.enter().append("g")
				.attr("class", "node")
				.attr("transform", d => `translate(${source.y0},${source.x0})`)
				.attr("fill-opacity", 0)
				.attr("stroke-opacity", 0)

			d3.selectAll('g.node')
				.attr("id", d => d.id)
				.on("click", this.clickNode)

			const nodeWidth = 200
			const nodeHeight = 50
			const nodeShape = nodeEnter.append('rect')
				.attr('x', -nodeWidth / 2)
				.attr('y', -nodeHeight / 2)
				.attr("rx", 15)
				.attr('width', nodeWidth)
				.attr('height', nodeHeight)
				.attr("stroke", d => { if (!d.data.color) { d.data.color = "#FFA500" } return d.data.color; })
				.attr("fill", d => d._children ? "#fff" : "#eee")
				.attr("opacity", 1)
				.attr("stroke-width", 5);


			nodeEnter.append("circle")
				.attr("class", "create")
				.style("visibility", "hidden")
				.attr("cx", nodeWidth / 2)
				.attr("r", 10)
				.attr("fill", "green")
				.attr("stroke", "#050")
				.attr("stroke-width", 3)
				.on("click", (event, d) => {
					event.stopPropagation();
					this.updateFromSky({
						name: "New Card",
						color: "#7f00ff",
						children: [],
						map_id: this.getCurrentMap.id,
						parent_id: this.getCurrentNode.data.id
					})
				});

			var image =	nodeEnter.append('image');
			this.nodeImg(image);

			nodeEnter.append("circle")
				.attr("class", "hide")
				.attr("cx", nodeWidth / 2)
				.style("visibility", d => d._children ? "visible" : "hidden")
				.attr("r", 10)
				.attr("fill", "#eee")
				.attr("stroke", "#ddd")
				.attr("stroke-width", 3)
				.on("click", (event, d) => {
					event.stopPropagation();
					d.children = d.children ? null : d._children;
					this.update(d);
					if (event && event.altKey) {
						setTimeout(() => {
							zoomToFit();
						}, duration + 100);
						//zoomToFit();
					}
				});

			var txt = nodeEnter.append("text")
			this.nodeText(txt)

			this.wrapText(nodeEnter);

			// Transition nodes to their new position.
			const nodeUpdate = node.merge(nodeEnter).transition(transition)
				.attr("transform", d => `translate(${d.y},${d.x})`)
				.attr("fill-opacity", 1)
				.attr("stroke-opacity", 1);

			// Transition exiting nodes to the parent's new position.
			const nodeExit = node.exit().transition(transition).remove()
				.attr("transform", d => `translate(${source.y},${source.x})`)
				.attr("fill-opacity", 0)
				.attr("stroke-opacity", 0);

			// Update the links…
			const link = gLink.selectAll("path")
				.data(links, d => d.target.id);

			// Enter any new links at the parent's previous position.
			const linkEnter = link.enter().append("path")
				.attr("d", d => {
					const o = { y: source.x0, x: source.y0 };
					return diagonal({ source: o, target: o });
				})

			// Transition links to their new position.
			link.merge(linkEnter).transition(transition)
				.attr("d", d => {
					return "M" + (d.target.y - nodeWidth / 2) + "," + d.target.x +
						"C" + (d.source.y + nodeWidth / 2 + 100) + "," + d.target.x +
						" " + (d.source.y + nodeWidth / 2 + 100) + "," + d.source.x +
						" " + (d.source.y) + "," + d.source.x;
				});

			// Transition exiting nodes to the parent's new position.
			link.exit().transition(transition).remove()
				.attr("d", d => {
					const o = { y: source.x0, x: source.y0 };
					return diagonal({ source: o, target: o });
				});

			// Stash the old positions for transition.
			root.eachBefore(d => {
				d.y0 = d.x;
				d.x0 = d.y;
			});

			return nodeEnter
		}
	}
}
</script>

<style scoped>
.canvas-style {
	cursor: crosshair;
	border: 1px solid black;
	display: block;
	margin: auto;
	box-shadow: 0 10px 8px -8px black;
}
</style>