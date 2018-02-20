//formu jsona ceviren fnksiyon
$.fn.serializeObject = function()
      {
          var o = {};
          var a = this.serializeArray();
          $.each(a, function() {
              if (o[this.name] !== undefined) {
                  if (!o[this.name].push) {
                      o[this.name] = [o[this.name]];
                  }
                  o[this.name].push(this.value || '');
              } else {
                  o[this.name] = this.value || '';
              }
          });
          return o;
      };



// loginsayfasında kayıt işlemi için
function KontrolEt(form) {
   
 } 
 
function KayıtEt(deger){

}

function Login(email,sifre) {
  
}

function Logout(deger) {

}

//diller alakalı
function getall_dil() {
  var tmp;
      $.ajax({
             url: 'http://127.0.0.1:8080/dil/all',
             data: {
                format: 'json'
             },
             error: function() {
                $('#info').html('<p>An error has occurred</p>');
             },
             dataType: 'json',
             success: function(data) {
              tmp=data;
              }
             },
             type: 'GET'
          });
  return tmp;
}

function create_dil(form) {
  var deger=JSON.stringify(form.serializeObject());
            
              $.ajax({
                     url: 'http://127.0.0.1:8080/dil/create',
                     data:deger,
                     error: function() {
                        alert('<p>Hata Oluştu</p>');
                     },
                     dataType: 'json',
                     success: function(re) {
                      var tmp=JSON.stringify(re);
                      return tmp;
                     },
                     type: 'POST'
                  });
}

function delete_dil(form) {
  
              //var deger= document.getElementById(ad);
              var s = form.options[form.selectedIndex].value;
              $('#sonuc').text(s);
              $.ajax({
                     url: 'http://127.0.0.1:8080/dil/delete',
                     data:s,
                     dataType: 'json',
                     error: function() {
                        alert('<p>Hata Oluştu</p>');
                     },
                     success: function(re) {
                      alert("basarılı");
                     },
                     type: 'DELETE'
                  });
}