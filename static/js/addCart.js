(function ($) {

    $(".addCart").on('click' ,function(e) {
        var id = $(this).val()
        var quantity = $("#quantity-" + id).val()
        console.log(id)
        window.location.href = "/product/add_cart:"+id+"&quantity=" + quantity
    })

})(jQuery);