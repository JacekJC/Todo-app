<script>
	import Task from "./task.svelte";
	import { onMount } from 'svelte';

	export let current_open_user = null;
	let tasks = null;

	let new_task_val;


	export async function load_tasks(user_name=null){
		tasks = null;
		console.log("loading tasks!")
		const task_fetch = await fetch("http://localhost:8081/user/"+user_name);
		const res = await task_fetch.json();
		console.log(res.Data)
		let json_data = JSON.parse(res.Data);
		console.log(json_data);
		tasks=json_data["Data"];
		console.log(tasks)
	}
	
	onMount(async () => {
		if(current_open_user!=null){
			load_tasks(current_open_user["name"]);
		}
	});

	async function add_new_task(){
		console.log(new_task_val, "new task added")
		if(new_task_val==undefined)
			return
		const task_add_fetch = await fetch("http://localhost:8081/add_task", {
			method:"POST",
			body: JSON.stringify({"Username":current_open_user["name"], "Taskname":new_task_val}),
		});
	}
	
</script>

<main>	
	<div class="top">
		{#if current_open_user!=null}
			<img this:use={load_tasks} src={current_open_user["img"]} alt="no img">
			<p> {current_open_user["name"]} </p>
		{/if}
	</div>
	<div>
		<div class="task-div">
			<p> Todays Tasks! </p>
			<div>
			{#if tasks != null}
				{#each tasks as task}
					<Task {task}/>
				{/each}
				<div class="new-task">
					<input placeholder="new task..." bind:value={new_task_val}>
					<div on:click={add_new_task}> + </div>
				</div>
			{/if}
			</div>
		</div>
	</div>

</main>

<style>

	.task-div div{
	}

	.task-div .new-task{
		background-color:white;
		margin:10px;
		width:50%;
		display:flex;
		padding:0;
	}

	.new-task input{
		font-size:1.1em;
		flex-grow:1;
		padding:10px;
		background-color: rgb(45,45,45);
		color:white;
		border:none;
	}

	.new-task div{
		user-select:none;
		background-color:rgb(40,40,40);
		font-size:2em;
		outline:solid 1px transparent;
		transition: all 0.1s;
	}
	
	.new-task div:hover{
		outline:solid 1px rgb(100,100,100);
		background-color:rgb(44,44,44);
	}


	.task-div{
		box-shadow: 0 0 10px -5px black;
		border-radius:10px;
		padding:5px;
		background-color:rgb(42,42,42);
		margin:10px;
	}

	main{
		display:grid;
		grid-template-rows:200px auto;
	}

	main .top{
		background-color: rgb(50,50,50);
		display:flex;
	}

	.top img{
		width:180px;
		margin:10px;
		height:180px;
		border-radius:20px;
	}

	.top p{
		user-select:none;
		font-size:2em;
		margin-top:2.4em;
		margin-left:1.4em;
	}


</style>
