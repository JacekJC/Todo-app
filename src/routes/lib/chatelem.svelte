
<script>
	export let chat;
</script>

<div class="chat-ob">
	<div>
		<img src={chat["user"]["img"]}>
		<p class="name" > {chat["user"]["name"]} </p>
		<p class="date"> 21:24 </p>
	</div>
	<p class="msg"> {chat["msg"]} </p>
</div>

<style>

	.msg{
		margin:0;
		padding-left:10px;
		padding-bottom:10px;
	}

	.name{
		flex-grow:1;
	}

	.date{
		right:0;
		padding-right:20px;
		color:rgb(150,150,150);
	}

	.chat-ob{
		user-select:none;
		margin:5px;
		background-color: rgb(60,60,60);
		border-radius:10px;
		color:white;
	}
	
	.chat-ob div{
		display:flex;
		color:white;
	}

	.chat-ob div img{
		margin:10px;
		width:30px;
		height:30px;
	}
</style>
