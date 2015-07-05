$.fn.tictacnine = function() {
  $(document).ready(function() {
    var container = $(this); 
    var game = function(container) {
      var internalGame = {
        reset : function() {
          container.find('.field-inner').each(function() {
            $(this).addClass('active');
            $(this).html('');  
          }); 
        }, 

        setContent : function(x, y, text) {
          var selector = '[data-pos-x="'+x+'"][data-pos-y="'+y+'"]';
          container.find(selector).attr('data-value', text); 
        }
      };
      
      // initially set all fields to active
      internalGame.reset(); 

      return internalGame; 
    }(container);

    console.log(game); 
    game.setContent(0, 0, 'x'); 
    game.setContent(1, 1, 'o'); 
  
  }); 
};
