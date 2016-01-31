(function() {
    var dom = {
        adminContent : $('.admin-content'),
        panelHeader : $('.me-panel-header')
    }

    var style = {
        init : function(){
            this.eventFn();
        },
        eventFn : function() {
            style.changeHeight();
            dom.panelHeader.bind('click',function() {
                console.log("jfjfj");
                style.changeHeight();
            });
        },
        changeHeight : function() {
            var height = dom.adminContent.height();
            if(height < 540) {
                dom.adminContent.height(540);
            } 
        }
    }
    style.init();
})();