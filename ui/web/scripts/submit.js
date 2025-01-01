
    function submitForm() 
        {
       var l= document.getElementById("inpunValue").value;
       var c= document.getElementById("cost").value;
       const myHeaders = new Headers();
       myHeaders.append("qq1","qq2")
      myHeaders.set("qq3","qqq4")

       
            fetch('http://localhost:8081/api/v1/calc',{ method:"POST", body: JSON.stringify({ liters: Number(l), cost:Number(c) }) })
            .then(response => response.json())
            .then(result => 
               { document.getElementById("zap").innerHTML="литров до полнойзаправки "+result.tofill,
                document.getElementById("cen").innerHTML="стоимость    запрвки "+result.totalcost  }          
            )
           ;


        }

