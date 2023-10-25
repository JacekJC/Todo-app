
<script>
	
	import {tick} from 'svelte'


	export let task;
	//export let task_name = "test task";
	//export let stage = 1;
	let click = null;

	function on_click(element){

		console.log(element.target);

		if(task["stage"]==0){
			task["stage"]=1;
		}
		else{
			task["stage"]=0;
		}
		update_state();

	}	

	function update_state(){
		if(click==null)
			return;

		switch(task["stage"]){
			case 0: 
				click.classList.add("full");
				click.classList.remove("empty");
				break;
			case 1: 
				click.classList.remove("full");
				click.classList.add("empty");
				break;
		}
	}

	function hi(){
		tick().then(()=>update_state());
	}

	function ping(){
		console.log("pinged user");
	}

	hi();

	update_state();

</script>

<div class="task"> 
	<p class="task_name"> {task["name"]} </p>
	<div class="delete">
		<p>X</p>
	</div>
	{#if task["stage"]==1}
		<div on:click={ping} class="ping"> <p>ping</p> </div>
	{:else}
		<div class="ping disabled" disabled> <p>ping</p> </div>
	{/if}
	<div class="click-div" on:click={on_click} bind:this={click} on:load={update_state}>
		<div hidden class="full empty"></div>
	</div>
</div>

<style>	
	.task{
		user-select:none;
		background-color: rgb(45,45,45);
		padding: 10px;
		margin-bottom:5px;
		display:flex;
	}

	.delete{
		display:flex;
		align_items:center;
		justify-content:center;
		margin-right:20px;
		border-radius:20px;
		width:30px;
		height:30px;
		background-color:rgb(160,70,60);
	}

	.delete p{
		transform:translate(5%,18.5%);
		margin:0;
	}

	.ping{
		background-color: rgb(55,55,55);
		margin:0;
		margin-right:10px;
		border-radius:2px;
		user-select:none;
	}
	.ping p{
		margin:0;
	}

	.ping:hover{
		transform:scale(105%);
	}

	.disabled{
		color:rgb(100,100,100);
		background-color:rgb(50,50,50);
	}

	.task .task_name{
		margin:0;
		transform:translate(0,5px);
		flex-grow:1;
	}

	.task .click-div{
		margin-top:4px;
		right:0;
		border: solid 1px white;
		border-radius:100%;
		width:20px;
		height:20px;
		position:relative;
		transition:all 0.1s;
	}

	.task .click-div:hover{
		/*border: solid 2px white;*/
		transform:scale(110%);
		text-align:center;
		justify-content:center;
	}

	.task div:hover{
	}

	.task .empty{
		background-color: transparent;
	}

	.task .full{
		background-color: white;
	}

</style>
