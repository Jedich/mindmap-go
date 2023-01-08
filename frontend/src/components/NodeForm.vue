<template>
	<div style="margin:10px;">
		<div style="display: block;">Name: <span class="input-group-addon limit"
				v-text="(nameLength - getCurrentNode.data.name.length)"></span>
			<textarea class="form-control" :value="getCurrentNode.data.name" @input="updateName" :maxlength="nameLength"
				rows="2" />
		</div>
		<div>Description:
			<textarea class="form-control" :value="getCurrentNode.data.text_data" @input="updateText" rows="5" />
		</div>
		<div>
			<div v-if="!imageError">
				<img :src="image.raw" class="uploading-image"
					style="max-width:calc(100% - 20px); max-height:auto; margin: 10px; border-radius: 5px;" />
			</div>
			<input type="file" ref="myFile" accept="image/*" @change=getUserFile>
		</div>
		<div class="alert alert-danger" role="alert" v-if="imageError">
			{{ imageError }}
		</div>
		<div v-if="image.raw && !imageError">
			<button style="margin-top:5px; width:100%;" class="btn btn-outline-success"
				v-on:click="uploadUserFile">Upload</button>
		</div>
		<div>Stroke color:
			<input type="color" style="width:100%" :value="getCurrentNode.data.color" @input="throttledColor" />
		</div>
		<div v-if="getCurrentNode.data.created">
			<button style="margin-top:5px; width:100%;" class="btn btn-outline-success"
				v-on:click="deselect">Create</button>
		</div>
		<div v-if="getCurrentNode.data.updated">
			<button style="margin-top:5px; width:100%;" class="btn btn-outline-success"
				v-on:click="deselect">Save</button>
		</div>
	</div>
</template>

<script>
import * as d3 from "d3";
import Canvas from './Canvas.vue';
import throttle from 'lodash/throttle';
import { mapActions, mapGetters } from "vuex";

const MAX_SIZE = 1000000;
const MAX_WIDTH = 1000;
const MAX_HEIGHT = 1000;

export default {
	computed: {
		...mapGetters("select", {
			getCurrentNode: "getCurrentNode",
		}),
	},
	components: {
		Canvas,
	},
	emits: ["updateSelection"],
	data() {
		return {
			nameLength: 45,
			imageError: '',
			image: {
				extension: '',
				size: '',
				height: '',
				width: '',
				raw: null,
				file: null
			},
		}
	},
	watch: {
		color: function () {
			this.throttledColor()
		},
		getCurrentNode: function () {
			this.imageError = '';
			this.image.raw = null;
			this.$refs.myFile.value = null;
		}
	},
	methods: {
		...mapActions("select", {
			select: "select",
			deselect: "deselect",
			uploadFile: "uploadFile",
			commitNode: "commitNode"
		}),
		updateName(event) {
			this.getCurrentNode.data.name = event.target.value
			if (!this.getCurrentNode.data.created) {
				this.getCurrentNode.data.updated = true;
			}
			this.$emit("updateSelection")
		},
		updateText(event) {
			this.getCurrentNode.data.text_data = event.target.value
			if (!this.getCurrentNode.data.created) {
				this.getCurrentNode.data.updated = true;
			}
		},
		throttledColor: throttle(function (event) {
			this.getCurrentNode.data.color = event.target.value
			this.getCurrentNode.s.select('rect').transition().duration(100).attr("stroke", d => d.data.color);
			if (!this.getCurrentNode.data.created) {
				this.getCurrentNode.data.updated = true;
			}
		}, 100),
		getUserFile(e) {
			this.imageError = '';
			this.previewImage = null;
			let file = this.$refs.myFile.files[0];

			if (!file || file.type.indexOf('image/') !== 0) return;
			
			this.image.size = file.size;
			if (this.image.size > MAX_SIZE) {
				this.imageError = `The image size (${this.image.size / 1000}KB) is too much (max is ${MAX_SIZE / 1000}KB).`;
				return;
			}
			this.image.extension = file.name.split('.').pop();
			let reader = new FileReader();

			reader.readAsDataURL(file);
			reader.onload = evt => {
				let img = new Image();
				img.onload = () => {
					this.image.width = img.width;
					this.image.height = img.height;
					if (this.image.width > MAX_WIDTH) {
						this.imageError = `The image width (${this.image.width}) is too much (max is ${MAX_WIDTH}).`;
						return;
					}
					if (this.image.height > MAX_HEIGHT) {
						return;
					}
				}
				if (this.imageError !== '') {
					return;
				}
				img.src = evt.target.result;
				this.image.raw = evt.target.result;
				this.image.file = file;
			}
			reader.onerror = evt => {
				console.error(evt);
			}
		},
		async uploadUserFile(e) {
			await this.uploadFile(this.image)
			this.$emit("updateSelection")
		}
	}
}
</script>