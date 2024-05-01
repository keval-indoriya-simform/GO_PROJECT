var emailDomainObj={}
var accountObj={}
var emailAccountObj={}
var emailAccountTypeObj={}
var customerLocationObj={}
let create_button = $("#submit-btn");
let customer_location_id = $("#customer_location_id");
let domain_registrar = $("#domain_registrar");
let account_number = $("#account_number");
let link_for_domain_admin = $("#link_for_domain_admin");
let account_name = $("#account_name");
let account_number_aliases = $("#account_number_aliases");
let account_user_name = $("#account_user_name");
let password = $("#password");
let pin = $("#pin");
let domain = $("#domain");
let a_record_1 = $("#a_record_1");
let a_record_2 = $("#a_record_2");
let a_record_3 = $("#a_record_3");
let a_record_4 = $("#a_record_4");
let email_hosting = $("#email_hosting");
let link_for_email_admin = $("#link_for_email_admin");
let email_account_number = $("#email_account_number");
let email_account_type_id = $("#email_account_type_id");
let email_user_name = $("#email_user_name");
let email_password = $("#email_password");
let email_pin = $("#email_pin");
let mx_record_1 = $("#mx_record_1");
let mx_record_2 = $("#mx_record_2");
let website_ip_or_alias=$("#website_ip_or_alias");
let web_mail_pop_imap=$("#web_mail_pop_imap");
let web_mail_exchange=$("#web_mail_exchange");
let phone_setting_note=$("#phone_setting_note");
let notes=$("#notes");

console.log("script loaded")




function emailAccountType(key,val){
    emailAccountTypeObj[key]=val
}
function customerLocation(key,val){
    customerLocationObj[key]=val

}
function createEmailDomain(key,val){
    emailDomainObj["customer_location"]=customerLocationObj
    emailAccountObj["email_account_type"]=emailAccountTypeObj
    emailDomainObj["account"]=accountObj

    emailDomainObj["email_account"]=emailAccountObj
}
$(document).ready(async function() {

    const customerLocationResponse = await fetch('http://192.168.49.2:31471/api/v1/customer-locations?select_column=customer_locations.customer_location_id,customer_locations.name');
    const customerLocationData = await customerLocationResponse.json();
    // console.log(customerLocationData)
    const emailAccountTypesResponse=await fetch('http://192.168.49.2:31471/api/v1/email-domains/email-account-types')
    const emailAccountTypesData=await emailAccountTypesResponse.json()
    console.log(emailAccountTypesData)


    options =""
    for (index = 0; index < customerLocationData["data"].length; index++) {
        options += `<option data-tokens="`+customerLocationData["data"][index]["name"]+`" value="`+ customerLocationData["data"][index]["customer_location_id"] + `">` + customerLocationData["data"][index]["name"] + `</option>`;
    }

    emailAccountTypeOptions=""
    for (index=0;index<emailAccountTypesData["data"].length;index++){
        emailAccountTypeOptions+=`<option data-tokens="`+emailAccountTypesData["data"][index]["email_account_type"]+`" value="`+ emailAccountTypesData["data"][index]["email_account_type_id"] + `">` + emailAccountTypesData["data"][index]["email_account_type"] + `</option>`;
    }

    console.log(emailDomainObj)
    $('.selectpicker').append(options).selectpicker('refresh');
    $('.selectpicker1').append(emailAccountTypeOptions).selectpicker('refresh');
    create_button.click(async function(){
        event.preventDefault()

        await createEmailDomainObject()
        console.log("hello")
        await postModel(
            JSON.stringify(emailDomainObj),
            "http://192.168.49.2:31471/api/v1/email-domains"
        )
    });


});

async function createEmailDomainObject(){
   emailDomainObj["customer_location_id"]=parseInt(customer_location_id.val())
    emailDomainObj["domain_registrar"]=domain_registrar.val()
    emailDomainObj["account"]={
       link_for_domain_admin:link_for_domain_admin.val(),
       account_number:account_number.val(),
       name:account_name.val(),
       account_number_aliases:account_number_aliases.val(),
       user_name:account_user_name.val()
    }
    emailDomainObj["email_account"]={
        email_hosting:email_hosting.val(),
        link_for_email_admin:link_for_email_admin.val(),
        account_number: email_account_number.val(),
        email_account_type_id:parseInt(email_account_type_id.val()),
        user_name:email_user_name.val(),
        password:email_password.val(),
        pin:email_pin.val()
    }
    emailDomainObj["password"]=password.val()
    emailDomainObj["pin"]=pin.val()
    emailDomainObj["domain"]=domain.val()
    emailDomainObj["a_record_1"]=a_record_1.val()
    emailDomainObj["a_record_2"]=a_record_2.val()
    emailDomainObj["a_record_3"]=a_record_3.val()
    emailDomainObj["a_record_4"]=a_record_4.val()
    emailDomainObj["mx_record_1"]=mx_record_1.val()
    emailDomainObj["mx_record_2"]=mx_record_2.val()
    emailDomainObj["website_ip_or_alias"]=website_ip_or_alias.val()
    emailDomainObj["web_mail_pop_imap"]=web_mail_pop_imap.val()
    emailDomainObj["web_mail_exchange"]=web_mail_exchange.val()
    emailDomainObj["phone_setting_note"]=phone_setting_note.val()
    emailDomainObj["notes"]=notes.val()
    emailDomainObj["created_by_user_id"] = parseInt($("#user_id").val())
    console.log(emailDomainObj["created_by_user_id"])
}
