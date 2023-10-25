<script>
	
	export let this_user = null;

	function clicked_func(element){
		clicked(this_user);
	}

	export let clicked;

</script>

<div on:click={clicked_func} class="person">
	<img src={this_user["img"]} alt="no img">
	<p> {this_user["name"]} </p>
</div>

<style>
	.person{
		margin:2px;
		padding:8px;
		display:flex;
		height:50px;
		/*box-shadow: 0 0 6px 0 black;*/
		outline: solid 1px transparent;
		border-radius: 2px;
		background-color:rgb(60,60,60);
		transition:all 0.05s;
	}

	.person:hover{
		outline: solid 1px rgb(220,220,255);
		border-radius: 2px;
		background-color:rgb(55,55,55);
		transform:scale(105%);
		border-radius:4px;
	}

	.person p{
		margin:0;
		margin-top:1em;
		padding:0;
		padding-left:10px;
		user-select:none;
	}
	
	.person img{
		width:50px;
		height:50px;
		border-radius:20px;
	}

</style>

