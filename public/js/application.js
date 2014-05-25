$('input[type="submit"]').on('click', function(e){
  e.preventDefault();
  $.ajax({
		type: 'POST',
		url: '/',
		data: $('form').serialize(),
		success: function(data){
			$('form').hide();
			$('#result').fadeIn();

			if (data.contained === true) {
				$('#success').fadeIn();
			} else {
				$('#failure').fadeIn();
			}
		},
	});
});

$('#reload').on('click', function(e){
	e.preventDefault();
	$('#success').hide();
	$('#failure').hide();
	$('#result').hide();
	$('form').fadeIn();
});