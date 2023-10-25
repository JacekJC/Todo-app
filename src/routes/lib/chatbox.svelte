<script>
	import Chatelem from './chatelem.svelte';

	let test = {
		user: {
			img:"https://i.pinimg.com/1200x/be/d0/ac/bed0ac8ce460608367137b40953db069.jpg",
			name:"shmaniel",
		},
		msg:"Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad ",
	};

</script>

<main hidden>
	<div class="chatbox">
		<div class="chat-div">
			<Chatelem chat={test} />
			<Chatelem chat={test} />
			<Chatelem chat={test} />
			<Chatelem chat={test} />
			<Chatelem chat={test} />
		</div>
		<div class="input-box">
			<input placeholder="message...">
			<div>send</div>
		</div>
	</div>
</main>

<style>

	.chat-div{
		overflow:hidden;
		overflow-y:scroll;
	}

	.chatbox{
		position:absolute;
		right:0;
		bottom:0;
		width:20%;
		height:50%;
		background-color:rgb(45,45,45);
		display:grid;
		grid-template-rows: auto 40px;
	}

	.input-box{
		display:flex;
	}

	.input-box input{
		flex-grow:1;
		color: white;
		background-color: rgb(35,35,35);
		border:none;
		padding-left:10px;
	}

	.input-box div{
		margin:0;
		color:white;
		background-color:rgb(50,50,50);
		padding:10px;
		padding-left:20px;
		padding-right:20px;
	}
</style>
