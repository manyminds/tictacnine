$.fn.tictacnine = function() {
  $(document).ready(function() {
    var container = $(this); 
    var game = function(container) {
      var internalGame = {
        reset : function() {
          container.find('.field-outer').each(function() {
            $(this).removeClass('active'); 
          }); 

          container.find('.field-inner').each(function() {
            $(this).addClass('active');
            $(this).html('');  
          }); 
        }, 

        setContent : function(x, y, text) {
          var selector = '[data-pos-x="'+x+'"][data-pos-y="'+y+'"]';
          field = container.find(selector); 
          if (!field) {
            throw "field not found";
          }

          field.attr('data-value', text); 
        }, 

        setField : function(x, y) {
          internalGame.reset();
          var selector = '[data-field-x="'+x+'"][data-field-y="'+y+'"]'; 
          field = container.find(selector); 
          if (!field) {
            throw "field not found";
          }

         field.addClass("active"); 
        }
      };
      
      // initially set all fields to active
      internalGame.reset(); 

      return internalGame; 
    }(container);

    game.setContent(0, 0, 'x'); 
    game.setContent(1, 1, 'o'); 
    game.setField(1, 1); 
  }); 
};
