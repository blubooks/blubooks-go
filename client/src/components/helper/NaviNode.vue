<template>    
    <li class="nav-item" :style="nodeMargin">

    <router-link :to="{ name: 'section', params: { id: node.id }}" @click="click()" class="nav-link" :class="{ active: node.id ==  $route.params.id }" aria-current="page">{{ node.title }}</router-link>
        <ul class="nav flex-column"  v-if="hasChildren" v-show="showChildren">
            <NaviNode
          v-for="child in node.children"
          :key="child.id"
          :node="child"
          :current="current"
          :spacing="spacing + 10"
      
        />              
        </ul> 
    </li>          
<!--
    <div :style="nodeMargin">
      <div show class="d-flex justify-content-between mb-1">
        {{ node.title }}
        <button
          v-if="hasChildren"
          :class="toggleChildrenIcon"
          @click="toggleChildren"
          @keypress="toggleChildren"
        />
      </div>
      <div v-if="hasChildren" v-show="showChildren">
        <NaviNode
          v-for="child in node.children"
          :key="child.id"
          :node="child"
          :spacing="spacing + 10"
        />
      </div>
    </div>
    -->
</template>
  
<script lang="ts">
import { toHandlers } from 'vue'

  export default {
    name: 'NaviNode',
    props: {
      node: {
        type: Object,
        required: true
      },
      current: {
        type: Object,
        required: true
      },      
      spacing: {
        type: Number,
        default: 0
      }
    },
    data() {
      return {
       // showChildren: false 
          showChildren: this.showChild()
      }
    },
    computed: {

      nodeMargin() {
        return {
          'margin-left': `${this.spacing}px`
        }
      },
      hasChildren() {
        if (this.$route.params.id == this.node.id) {
            this.showChildren = true;
        }
 
        const { children } = this.node
        return children && children.length > 0
      },
      toggleChildrenIcon() {

        return this.showChildren ? 'fas fa-angle-down' : 'fas fa-angle-right'
      }
    },
    methods: {
        showChild() {
            if ( this.showChildren ) {
                return false;
            }
            if (this.current.ids.includes( this.node.id)){
               return true;
            }
        return false
    },      
    click() {
        this.showChildren = !this.showChildren
    } ,
      show() {
        this.showChildren = true   
      //  this.$emit("show", "");    

      },        
      toggleChildren() {
        this.showChildren = !this.showChildren
      }
    }
  }
  </script>