$.fn.tictacnine = function(options) {
  $(document).ready(function() {
    if (typeof options == 'object') {
      var url = options.endpoint;
      if (!url) {
        alert("endpoint not specified");
      }
    
      var userColor = options.userColor;
      if (userColor != 'X') {
        userColor = 'O';
      } else {
        userColor = 'X';
      }

      var aiStrength = options.aiStrength; 
    
      if (! Number.isInteger(aiStrength)) {
        aiStrength = 5;
      }

      if (aiStrength >= 10 || aiStrength <= 1) {
        aiStrength = 6;
      }

      console.log("Settings: ");
      console.log({
        'aiStrength' : aiStrength, 
        'userColor' : userColor,
        'endpoint' : url
      });

      $.get("/ai/init", function(data) {
      })
}

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

        CalcFieldWon : function(x, y) {
          var selector = '[data-field-x="'+x+'"][data-field-y="'+y+'"]'; 
          field = container.find(selector); 
          if (!field) {
            throw "field not found";
          }

          if (field.find('.won-x').length > 0 || field.find('.won-o').length > 0) {
            return; 
          }

          var validationLines = [
            [0, 4, 8], 
            [0, 1, 2], 
            [0, 3, 6], 
            [1, 4, 7], 
            [2, 5, 8], 
            [2, 4, 6],
            [3, 4, 5], 
            [6, 7, 8]
          ]; 

          var board = new Array(); 

          field.find('.field-inner').each(function() {
            var value = $(this).attr('data-value'); 
            var x = $(this).attr('data-pos-x') % 3, 
                y = $(this).attr('data-pos-y') % 3;  
         
            var pos = x+(y*3);

            if (value) {  
             if (value == 'x') {
                board[pos] = 0;  
              } else {
                board[pos] = 1;  
              }
            } else {
              board[pos] = 9; 
            }
          });

          for (var row = 0; row < validationLines.length; row++) {
            var fields = validationLines[row]; 
            var sum = board[fields[0]] + board[fields[1]] + board[fields[2]]; 

            if (sum == 0) {
              field.find('tbody').before('<div class="won-x"></div>'); 
            } else if(sum == 3) {
              field.find('tbody').before('<div class="won-o"></div>'); 
            } 
          }
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
              $.post("/ai/putStone", {'x' : x, 'y': y, 'move': 'x'}, function(data) {
                  b.SetContent(x,y, 'x');              
                  activePlayer = playerTwo;
                  $.get("/ai/getStone", function(data) {
                    b.SetContent(data.x, data.y, 'o'); 
                    b.SetField(data.x %3 , data.y%3, 'o'); 
                    activePlayer = playerOne;
                  })
              })
            } else {
              alert("ai turn"); 
              return;
              b.SetContent(x,y, 'o');              
              activePlayer = playerOne; 
            }
            
            b.CalcFieldWon(fieldX, fieldY);  

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
