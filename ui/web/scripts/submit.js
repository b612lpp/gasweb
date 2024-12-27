
    function submitForm() 
        {
       var l= document.getElementById("inpunValue").value;
       var c= document.getElementById("cost").value;
       
            fetch('http://localhost:8081/api/v1/calc',{method:"POST", body: JSON.stringify({ liters: Number(l), cost:Number(c) }) })
            .then(response => response.json())
            .then(result => 
               { document.getElementById("zap").innerHTML="литров до полно    запрвки "+result.tofill,
                document.getElementById("cen").innerHTML="стоимость    запрвки "+result.totalcost  }          
            )
           ;


        }

