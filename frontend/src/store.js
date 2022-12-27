import { reactive } from 'vue'

export const store = reactive({
  selectedNode: null,
  putNode(node) {
    this.selectedNode = node
  }
})