<script>
	export let elements;
	let self;

	let dragging = false
	let pos_store = null;
	let before_size = null;

	let full_before_size = null;

	let mouse_move_listen = (e) => {
		console.log(before_size)
		before_size+=e.clientX-pos_store;
		pos_store = e.clientX;
		console.log(before_size+"px")
		elements[0].style.width = before_size+"px"
		if(elements[1]!=null){
			elements[1].style.width = full_before_size-before_size+"px";
		}
		console.log(elements[0].style.width)
	}

	function drag_start(e){	
		dragging=true
		console.log("drag-start", e)
		pos_store = e.clientX;
		before_size = parseInt(elements[0].style.width);
		if(elements[1]!=null){
			full_before_size = before_size + parseInt(elements[1].style.width)
		}
		console.log(before_size);
		window.addEventListener("mousemove", mouse_move_listen);
		window.addEventListener("mouseup", drag_end);
		//self.style.width="8px"
	}
	
	function drag_end(){
		dragging=false
		pos_store = null;
		console.log("drag-end")
		if(elements[1]!=null){
			full_before_size=null;
		}
		window.removeEventListener("mousemove", mouse_move_listen);
		window.removeEventListener("mouseup", drag_end);
		//self.style.width="4px"
	}	

	//function mouse_move(e){
	//	if(!dragging)
	//		return
	//	console.log("mouse-move!")
	//	console.log(e)
	//	before_size+=e.clientX-pos_store;
	//	pos_store = e.clientX;
	//	element.style.width = before_size+"px;"
	//}

</script>

<div 
	class="dragbar"
	on:mousedown={drag_start}
	draggable=false
	bind:this={self}
>
	
</div>

<style>
	.dragbar{
		width:4px;
		background-color:rgb(20,20,20);
		margin:4px;
	}

	.dragbar:hover{
		cursor: col-resize;
	}

</style>

