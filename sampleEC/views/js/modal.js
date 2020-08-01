$(function () {
    $('.js-modal-open').on('click', function () {
        var now = new Date();
        var wd = ['日', '月', '火', '水', '木', '金', '土'];
        var selectValue = 1;

        for (var i = 0; i < 7; i++) {
            var m = now.getMonth() + 1;
            var d = now.getDate();
            var w = wd[now.getDay()];
            now.setDate(now.getDate() + 1);
            $('.dateSlide').append($('<option>').html(m + "月" + d + "日" + "(" + w + ")").val(selectValue));
            selectValue++;
        }

        $('.js-modal').fadeIn();
        return false;
    });
    $('.js-modal-close').on('click', function () {
        $('.js-modal').fadeOut();
        return false;
    });
});