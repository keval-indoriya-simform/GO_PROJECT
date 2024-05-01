$( document ).ready(function() {
    $('.customer_name , .customer_location').selectpicker();
    $('.customer_name , .customer_location').selectpicker('val', 'Mustard');

    $(".customer_name input").on('change keyup paste',function(){
        console.log("asdasd")
    });
    let customers;
    func1().then(data=>{
        customers=data.data
        console.log(customers)
        for (var index = 0; index < customers.length; index++) {

            $('.customer_name ul').append('<li class="selected active"><a role="option" class="dropdown-item active selected" id="bs-select-1-0" tabindex="0" aria-setsize="1" aria-posinset="1" aria-selected="true"><span class="text">Simform</span></a></li>');

        }
    } );



});


const func1=async ()=>{
    let options = {
        method: 'GET',
        headers: {
            'Content-Type': 'application/json;charset=utf-8'
        },
    }
    const req=await fetch("http://192.168.49.2:31471/api/v1/customers",options);
    return await req.json();
}