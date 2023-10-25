<script>
	import Listperson from './listperson.svelte'
	
	//let temp = {"name":"kyle UwU"};

	let temp = {name:"kyle UwU :3", img:"null"};
	let tempa = {name:"Dan teehee :S", img:"null"};
	let tempb = {name:"Jakook x)", img:"null"};

	export let me_user = null;

	export let side_callbacks;

</script>

<main>
	<div class="category" >
		<p> General </p>
		{#if me_user!=null}
			<Listperson this_user={me_user} clicked={side_callbacks[0]}/>
		{/if}
		<div hidden class="chat-button" >
			Chat
		</div>
	</div>
	<div class="category" >
		<p> Online Users </p>
	</div>
	<div class="category">
		<p> Stuff </p>
	</div>
</main>

<style>
	.chat-button{
		text-align:center;
		height:50px;
		verical-align:center;
		background-color:rgb(60,60,60);
		border-radius:5px;
	}

	main{
		display:flex;
		width:100%;
		flex-flow:column;
		color:white;
		height:100%;
	}

	.category{
		display:flex;
		flex-flow:column;
		gap:4px;
		padding:2px;
		background-color:rgb(48,48,48);
		margin:4px;
		border-radius:2px;
	}

	.category p{
		text-align:center;
		margin:4px;
		user-select:none;
	}

	

</style>
