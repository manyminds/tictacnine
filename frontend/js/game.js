$.fn.tictacnine = function() {
  $(document).ready(function() {
    var container = $(this); 
    var board = function(container) {
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

        SetContent : function(x, y, text) {
          var selector = '[data-pos-x="'+x+'"][data-pos-y="'+y+'"]';
          field = container.find(selector); 
          if (!field) {
            throw "field not found";
          }

          field.attr('data-value', text); 
        }, 

        GetContent : function(x, y) {
           var selector = '[data-pos-x="'+x+'"][data-pos-y="'+y+'"]';
          field = container.find(selector); 
          if (!field) {
            throw "field not found";
          }

          return field.attr('data-value');
        },

        SetField : function(x, y) {
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
    };

    var game = function(board, container) {
      var playerOne = "player 1", 
          playerTwo = "player 2", 
          activePlayer; 
      
      return {
        init : function() {
          activePlayer = playerOne; 
          var b = board(container);

          container.find('.field-inner').click (function() {
            var x = $(this).data('pos-x'); 
            var y = $(this).data('pos-y');
            if (b.GetContent(x, y) != undefined) {
              alert("Already used."); 
              return; 
            }

            if (activePlayer === playerOne) {
              b.SetContent(x,y, 'x');              
              activePlayer = playerTwo; 
            } else {
              b.SetContent(x,y, 'o');              
              activePlayer = playerOne; 
            }
          }); 
        } 
      }; 
    }(board, container); 
  
    game.init(); 
  }); 
};
