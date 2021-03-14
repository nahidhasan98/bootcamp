$(document).ready(function () {
    //alert("loaded succesfully");

    $('#reqForm').submit(function () {
        //alert('submit btn pressed');

        $.ajax({
            async: true,
            type: "POST",
            url: "/request",
            data: $('#reqForm').serialize(),
            error: function (err, stsco) {
                alert(e, stsco);
            }
        }).done(function (rdata) {
            //alert(rdata);
            location.reload();
        });
        return false;
    });
});