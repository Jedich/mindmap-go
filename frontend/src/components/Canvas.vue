<template>
	<svg class="canvas"></svg>
</template>

<script>
import * as d3 from "d3";

export default {
	mounted() {
		const width = 800;
		const data = {
			name: "text text text text text text text text text",
			children: [
				{
					name: "123490\n1123456711234 567890112345678901",
					children: [{
						name: "b",
						color: "green",
						children: []
					},
					{
						name: "c",
						children: []
					},
					{
						name: "c",
						children: []
					}]
				},
				{
					name: "c",
					children: [{
						name: "b",
						children: []
					},
					{
						name: "c",
						children: []
					}]
				}
			]
		}
		var diagonal = d3.linkHorizontal().x(d => d.y).y(d => d.x)
		const root = d3.hierarchy(data);
		var dy = width / data.children.length;
		var dx = 120;

		root.x0 = 0;
		root.y0 = dy / 2;
		root.descendants().forEach((d, i) => {
			d.id = i;
			d._children = d.children;
			if (d.depth) d.children = null;
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

		const zoomBehaviours = d3.zoom()
			.scaleExtent([0.05, 3])
			.on('zoom', (event) => g.attr('transform', event.transform))
			.filter((event) => (event.button === 1) || event.type === 'wheel');

		svg.call(zoomBehaviours);

		setTimeout(() => zoomBehaviours.translateTo(svg, 0, 0), 100);

		function update(source) {
			const duration = d3.event && d3.event.altKey ? 2500 : 250;
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
				.data(nodes, d => d.id);

			// Enter any new nodes at the parent's previous position.
			const nodeEnter = node.enter().append("g")
				.attr("transform", d => `translate(${source.y0},${source.x0})`)
				.attr("fill-opacity", 0)
				.attr("stroke-opacity", 0)
				.on("click", function (event, d) {
					d.children = d.children ? null : d._children;
					update(d);
					if (event && event.altKey) {
						setTimeout(() => {
							zoomToFit();
						}, duration + 100);
						//zoomToFit();
					}
				});

			const nodeWidth = 150
			const nodeHeight = 50
			const nodeShape = nodeEnter.append('rect')
				.attr('x', -nodeWidth / 2)
				.attr('y', -nodeHeight / 2)
				.attr("rx", 15)
				.attr('width', nodeWidth)
				.attr('height', nodeHeight)
				.attr("stroke", d => d.color != null ? d.color : "red")
				.attr("fill", d => d._children ? "#fff" : "#eee")
				.attr("opacity", 1)
				.attr("stroke-width", 5);

			nodeEnter.append("circle")
				.attr("cx", nodeWidth / 2)
				.attr("r", 10)
				.attr("fill", "#eee")
				.attr("stroke", "#ddd")
				.attr("stroke-width", 3);

			nodeEnter.append("text")
				.attr("x", 0)
				.attr("y", d => { d.data.wrappedText = wrapRecu(d.data.name, 13); return -15 - (d.data.wrappedText.length - 1) * 10 })
				.attr("dy", "0em")
				// .clone(true)
				// .lower()
				.attr("font-size", 20)
				.text(d => d.data.wrappedText.join("/"));

			wrap(nodeEnter.selectAll('text'), 5);

			function wrapRecu(text, limit, resultArr = []) {
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
					return wrapRecu(remainder, limit, resultArr);
				} else {
					resultArr.push(text)
				}
				if (resultArr.length === 0) {
					return [text]
				}
				return resultArr;
			}

			function wrap(text, width) {
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
					words.forEach(word => {
						tspan = text.append("tspan").attr("x", x).attr("text-anchor", "middle").attr("y", y).attr("dy", ++lineNumber * lineHeight + dy + "em").text(word);
					});
					// find corresponding rect and reszie
					var h = 50 + ((lineNumber - 1) * 17)
					d3.select(this.parentNode.children[0]).attr('height', h).attr('y', -h / 2);

				});
			}

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
					// return "M" + d.target.y + "," + d.target.x +
					// 	"C" + (d.source.y + 200) + "," + d.target.x +
					// 	" " + (d.source.y + 100) + "," + d.source.x +
					// 	" " + d.source.y + "," + d.source.x;
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
		}

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

		// var height = 800;

		// var scale = d3.scaleLinear()
		// 	.domain([0, 100])
		// 	.range([0, width]);

		// var x = d3.scaleBand().rangeRound([0, width]).padding(0.4);

		// // Axis
		// var axis = d3.axisTop()
		// 	.scale(scale);

		// var axis2 = d3.axisLeft()
		// 	.scale(x);

		// svg.append("g")
		// 	.call(axis);

		// svg.append("g")
		// 	.call(axis2);

		// // Gridline
		// var gridlines = d3.axisTop()
		// 	.tickFormat("")
		// 	.tickSize(-height)
		// 	.scale(scale);

		// svg.append("g")
		// 	.attr("class", "grid")
		// 	.call(gridlines);

		// var gridlines2 = d3.axisLeft()
		// 	.tickFormat("")
		// 	.tickSize(-height)
		// 	.scale(scale);

		// svg.append("g")
		// 	.attr("class", "grid")
		// 	.call(gridlines2);

		update(root);

		// root.select('rect')
		// 	.attr('x', -nodeWidth * 1.1 / 2)
		// 	.attr('y', -nodeHeight * 1.1 / 2)
		// 	.attr("rx", 15)
		// 	.attr('width', nodeWidth * 4)
		// 	.attr('height', nodeHeight * 4)
		// 	.attr("stroke", "red")
		// 	.attr("fill", "#fff")
		// 	.attr("opacity", 1)
		// 	.attr("stroke-width", 5);

		//setTimeout(() => { zoomToFit();}, 5000);

		return svg.node();
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