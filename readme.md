# Go Language REST API
> ## USER API (./rest/user.go)
>> ### GET APIs 
>>> 1. __/userall__   ->  this api is used for the admin pourpose
 __localhost:8080/userall__
>>> 2.  __/user/{userID}__ -> this api use by the admin and check the details of the particular customer
__localhost:8080/user/User000__
>> ### POST APIs
>>>1. __/Login__ -> is used for login the user and get back the details 
__localhost:8080/login?email=xxxx@xxxx.com&pass=xxxxxxx__
>>> 2. __/user__ -> Register a new user  __localhost:8080/user__
        this json need to be share as USERs type object <br/>
        {<br/>
               "name":"ABC",<br/>
              "email":"ABC@gmail.com",<br/>
              "password":"ABCXXXX",<br/>
              "role":"customer/Admin"<br/>
        }
>> ### PUT API
>>> 1. __/user__  as put in request thet request helps to update details of the user
__localhost:8080/user/userId__ <br/>
        this json need to be share as Updated User Detials type object<br/>
        {<br/>
               "name":"ABC",<br/>
              "email":"ABC@gmail.com",<br/>
              "password":"ABCXXXX",<br/>
              "role":"customer/Admin"<br/>
        }
>> ### DELETE API 
>>> 1. __/user__ as Deleted the userId  
__localhost:8080/user/UserId__

> ## CATEGORY API
>> ### GET API
>>> __ALL CATEGORY__ : for admin and also for viewer we need to get all the category 
__localhost:8080/category__
>> ### POST API
>>> __CREATE CATEGORY__: for creating the new category we need to pass with __localhost:8080/category__ <br/>
        {
        "name":"Mobiles"
        }
>> ### PUT API
>>> __Update category Name__ : For rename the category name
        __localhost:8080/category/categoryId__
         <br/> {
        "name":"Moto Mobiles"
        }
>>  ### DELETE API
>>> __Delete Category__: for Deleting the category 
__localhost:8080/category/categoryId__

> ## Product API
>> ### GET API
>>> __View all products__: this api can hit for the index.html in which we able to see every product
__localhost:8080/product__<br/>
>>> __View single Product__: This api gives the details of the particular product
__localhost:8080/product/productId__
>> ### POST API
>>> NEW Product:- This api hit by the admin and enter the product details and available for the customer
__localhost:8080/product__<br/>
    {
    "name":"moto one fusion plus",
    "price":"19000",
    "description":";lsdakfjsadljflkadsjlfadsjkladslkfadsjfewoflksadjoir",
    "stock":100,
    "category":{
        "name":"Mobiles"
    }
    }
>> ### PUT API
>>> update product details : Admin can update the details of the product
__localhost:8080/product/productId__
<br/>
    {
    "name":"moto one fusion plus",
    "price":"19000",
    "description":";lsdakfjsadljflkadsjlfadsjkladslkfadsjfewoflksadjoir",
    "stock":100,
    "category":{
        "name":"Mobiles"
    }
    }
>> ### Delete Product 
>>> Delete a particular product from the productList that process done by the Admin
__localhost:8080/product/productId__


> ### Two more flow are same like those Orders and cart