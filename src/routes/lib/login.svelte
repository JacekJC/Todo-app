<script>
</script>

<main>
	<div class="main-div">
		<div class="login-box">
			<div class="title"><p> Cooptodo </p></div>
			<form class="login-form">
				<div><p> Username: </p><input></div>
				<div><p> Password: </p><input></div>
				<button> login </button>
			</form>
		</div>
	</div>
</main>

<style>
	.main-div{
		width:100%;
		height:100%;
		position:absolute;
		top:0;
		left:0;
		background-color:rgb(40,40,40);
		color:white;
		display:flex;
		text-align:center;
		justify-content:center;
		flex-justify:center;
		align-items: center;
	}

	.login-box{
		box-shadow:0 0 40px -20px black;
		margin:10px;
		padding:20px;
		width:400px;
		background-color: rgb(50,50,50);
		border-radius:10px;
	}

	.title{
		padding:20px;
		border-radius:10px;
		font-size:2em;
		background-color:rgb(60,60,60);
	}

	p{
		margin:0;
	}

	.login-form{
		margin-top:10px;
		display:flex;
		flex-flow:column;
		gap:10px;
	}

	.login-form div{
		display:flex;	
	}

	.login-form div p{
		width:20%;
		padding:7px;
		background-color:rgb(45,45,45);
	}

	.login-form input{
		flex-grow:1;
		height:30px;
		border:none;
		border-radius:5px;
		border-top-left-radius:0;
		border-bottom-left-radius:0;
		background-color:rgb(30,30,30);
		color:white;
		padding-left:10px;
	}

	.login-form button{
		border-radius:4px;
		height:30px;
		border:none;
		background-color: rgb(30,30,30);
		outline:solid 1px rgba(1,1,1,0);
		color:white;
		transition:all 0.1s;
	}

	.login-form button:hover{
		transform:scale(101%);
		outline:solid 1px rgba(255,255,255,0.4);
	}
</style>
