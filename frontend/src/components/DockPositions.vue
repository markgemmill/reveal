<template>
<div class="dock-options">
   <div class="box first" :class="isSelected('center')" @click="selectPosition('center')">
        <div class="position center"></div>
   </div> 
   <div class="box" :class="isSelected('ul')" @click="selectPosition('ul')">
        <div class="position upper-left"></div>
   </div> 
   <div class="box" :class="isSelected('ur')" @click="selectPosition('ur')">
        <div class="position upper-right"></div>
   </div> 
   <div class="box" :class="isSelected('ll')" @click="selectPosition('ll')">
        <div class="position lower-left"></div>
   </div> 
   <div class="box" :class="isSelected('lr')" @click="selectPosition('lr')">
        <div class="position lower-right"></div>
   </div> 
</div>
</template>

<script setup lang="ts">
import { watch } from "vue"
import type { DockPosition } from './types';
const position = defineModel<DockPosition>('position', {required: true})

const selectPosition = (thisPosition: DockPosition) => {
   console.debug(`selecting: ${thisPosition}`) 
   position.value = thisPosition
   console.debug(`position: ${position.value}`) 
}

const isSelected = (thisPosition: DockPosition) => {
    return {
        "selected": thisPosition === position.value
    } 
}

watch(position, (value) => {
    console.debug(`new position: ${position.value}`) 
})

</script>

<style scoped>
.dock-options {
    position: relative;
    display: flex;
    flex-direction: row;
    padding: 0px;
    /* border: 1px solid #CCC; */
}
.box {
    position: relative;
    display: flex;
    height: 40px;
    width: 40px;
    min-height: 40px;
    min-width: 40px;
    max-height: 40px;
    max-width: 40px;
    border: 1px solid #CCC;
    border-radius: 3px;
    margin-left: 10px;
}
.box:hover {
    cursor: pointer;
    border-color: #888;
}
.box:hover .position {
    background-color: #888;
}

.selected {
    border-color: black;
}
.selected .position {
    background-color: black;
}

.first {
    margin-left: 0;
}
.position {
    position: absolute;
    height: 20px;
    width: 20px;
    min-height: 20px;
    min-width: 20px;
    max-height: 20px;
    max-width: 20px;
    background-color: #CCC;
}
.center {
    top: 10px;
    left: 10px;
}
.upper-left {
    top: 0;
    left: 0;
    border-top-left-radius: 3px;
}
.upper-right {
    top: 0;
    right: 0;
    border-top-right-radius: 3px;
}
.lower-left {
    left: 0;
    bottom: 0;
    border-bottom-left-radius: 3px;
}
.lower-right {
    right: 0;
    bottom: 0;
    border-bottom-right-radius: 3px;
}
</style>