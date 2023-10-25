<script>

	let ip = "localhost"

	let user = null;

	let sock = null;

	import { onMount } from 'svelte';

	import Sidebar from './lib/sidebar.svelte';
	import Userview from './lib/userview.svelte';
	import Chatbox from './lib/chatbox.svelte';
	import Login from './lib/login.svelte';
	import Dragbar from './lib/dragbar.svelte';

	let me_user = null;
	
	let user_view = null;
	
	let image_link = "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcROxxkp8eXBn6NjQdZXFse9KJ8GqSpWZOsubA&usqp=CAU";

	let current_open_user = {"name": "Kyre UwU", "image": image_link};

	async function open_user(user){
		console.log("opening user! "+user["name"]);
		current_open_user = null;
		current_open_user = user;
		user_view.load_tasks(user["name"]);
	}

	async function login(username, password){
		console.log("posting")
		const response = await fetch("http://"+ip+":8081/login", {
			method: "POST",
			body: JSON.stringify({"Username":username, "Password":password})
		})
		let jresult = await response.json();
		jresult = JSON.parse(jresult["Data"])
		console.log(jresult);
		user = jresult;
		me_user = jresult;
	}


	let side_callbacks = [open_user];

	//var ping_sound = new Audio('ping.mp3');

	function play_sound(sound){
		audio.play();
	}


	function socket_connect(){
		console.log("connecting to server!")
		sock = new WebSocket("ws://"+ip+":8081/ws")
		sock.onclose = () => (console.log("socket closed"));
		sock.onmessage = (msg) => {
			console.log("rec: ",  msg);
		};
		console.log(sock)
	}

	async function get_profile(user_name){
		let t = await fetch("http://"+ip+":8081/profile/user/" + user_name);
		let res = await t.json();
		console.log(res);
		me_user = JSON.parse(res["Data"]);
		console.log(me_user)
	}

	onMount(() => {
		socket_connect();
		login("Chuube", "123funnyman")
		//get_profile("Chuube");
	})


	let friends_list_elem = null;
	function friends_list_toggle(){
		if(friends_list_elem==null)
			return;
		if(friends_list_elem.style!="width:300px;"){
			friends_list_elem.style="width:300px;"
		}else{
			friends_list_elem.style="width:0px;"
		}
	}

	let side_bar_elem;
	let user_view_elem;

</script>


<main on:load{socket_connect}>
	<head>
		<link rel="preconnect" href="https://fonts.googleapis.com">
		<link rel="preconnect" href="https://fonts.gstatic.com" crossorigin>
		<link href="https://fonts.googleapis.com/css2?family=Montserrat&display=swap" rel="stylesheet">
	</head>
	{#if user==null}
		<Login />
	{:else}
		<div class="main-div">
			<div class="top-bar">
				{#if me_user!=null}
					<img src={me_user["img"]}>
					<p> {me_user["name"]} </p>
				{/if}
			</div>
			<div class="sub-div">
				<div style="width:300px;" class="side-bar" bind:this={side_bar_elem}>
					<Sidebar {side_callbacks} {me_user}/>
				</div>
				<Dragbar elements={[side_bar_elem, user_view_elem]} />
				<div style="width:1200px;" bind:this={user_view_elem} class="user-view">
					{#if current_open_user!=null}
						<Userview bind:this={user_view} {current_open_user} />
					{/if}
				</div>
				<Dragbar elements={[user_view_elem, null]} />
				<div bind:this={friends_list_elem} class="friends-list">
				</div>
			</div>
		</div>
		<Chatbox />
	{/if}
</main>

<style>

	.side-bar{
		display:flex;
	}

	.top-bar{
		padding-top:12.5px;
		padding-left:15px;
		display:flex;
		gap:20px;
	}

	.top-bar img{
		border-radius:10%;
		height:70px;
		width:70px;
		object-fit: cover;
	}

	.top-bar p{
		margin:0.5em;

		font-size:2em;
	}

	.main-div div{
		font-family: 'Montserrat', sans-serif;
		/*border: solid 2px rgb(60,60,60);*/
		box-shadow: 0 0 10px -2px rgba(0,0,0,0.6);
		border-radius:10px;
		margin:2px;
	}
	.main-div{
		background-color:rgb(40,40,40);
		position:absolute;
		top:0;
		left:0;
		width:100%;
		height:100%;
		color:white;
		display:grid;
		grid-template-rows: 100px auto;
	}

	.sub-div{
		/*display:grid;
		grid-template-columns: 300px auto 300px;*/
		display:flex;
	}

	.user-view{
	}

	.friends-list{
		position:relative;
	}

	.friends-list button{
		width:30px;
		height:30px;
		position:absolute;
		top:0;
		left:0;
	}
</style>
