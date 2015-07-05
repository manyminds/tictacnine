$.fn.tictacnine = function() {
  $(document).ready(function() {
    var container = $(this); 
    var board = function(container) {
      var internalGame = {
        reset : function() {
          container.find('.field-outer').each(function() {
            $(this).addClass('active');
          }); 

          container.find('.field-inner').each(function() {
            $(this).addClass('active');
            $(this).html('');  
          }); 
        }, 

        showError : function(msg) {
          //TODO display clean errors, for now:
          alert(msg); 
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

        SetField : function(x, y, clear) {
          if (clear) {
             container.find('.field-outer').each(function() {
               $(this).find('.field-inner').each(function() {
                $(this).removeClass('active'); 
               }); 

               $(this).removeClass('active'); 
            }); 
          }
          
          var selector = '[data-field-x="'+x+'"][data-field-y="'+y+'"]'; 
          field = container.find(selector); 
          if (!field) {
            throw "field not found";
          }
         
         field.addClass("active");
         field.find('.field-inner').each(function() {
          $(this).addClass('active'); 
         });  
        }, 

        IsActiveField : function(x, y) {
          var selector = '[data-field-x="'+x+'"][data-field-y="'+y+'"]'; 
          field = container.find(selector); 
          if (!field) {
            throw "field not found";
          }

          return field.hasClass('active'); 
        }, 
  
       IsFullField : function(x, y) {
          var selector = '[data-field-x="'+x+'"][data-field-y="'+y+'"]'; 
          field = container.find(selector); 
          if (!field) {
            throw "field not found";
          }

          var numElems = 0; 
          field.find('.field-inner').each(function() {
            if ($(this).attr('data-value')) {
              numElems++; 
            }
          }); 

          return numElems == 9;  
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
              b.showError("Already used."); 
              return; 
            }

            var fieldX = Math.floor(x/3); 
            var fieldY = Math.floor(y/3); 

            if (! b.IsActiveField(fieldX, fieldY)) {
              b.showError("only the active field is allowed.");
              return 
            }

            if (activePlayer === playerOne) {
              b.SetContent(x,y, 'x');              
              activePlayer = playerTwo; 
            } else {
              b.SetContent(x,y, 'o');              
              activePlayer = playerOne; 
            }

            var targetFieldX = x%3; 
            var targetFieldY = y%3; 
            if (! b.IsFullField(targetFieldX, targetFieldY)) {
              b.SetField(targetFieldX, targetFieldY, true); 
            } else {
              for (var fx = 0; fx < 3; fx++) {
                for(var fy = 0; fy < 3; fy++) {
                  if (! b.IsFullField(fx, fy)) {
                    b.SetField(fx, fy, false);  
                  }
                }
              }
            }
          }); 
        } 
      }; 
    }(board, container); 
  
    game.init(); 
  }); 
};
