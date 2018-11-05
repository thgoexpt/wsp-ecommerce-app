(function ($) {

    $(".addCart").on('click' ,function(e) {
        var id = $(this).val()
        console.log("ID: " + id)
        var quantity = $("#quantity-" + id).val()
        console.log("Qty: " + quantity)
        if ( quantity === undefined) {
            console.log("under the roof")
            window.location.href = "/product/add_cart:"+id+"&quantity=1/"
        } else {
            window.location.href = "/product/add_cart:"+id+"&quantity=" + quantity + "/"
        }
    })

})(jQuery);