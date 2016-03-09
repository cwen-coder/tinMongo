(function() {
    var dom = {
        dbTransferAll : $('#dbTransferAll'),
        dbCollections : $("form div:first input[name='collection']")
    }

    var dbTransfer = {
        init : function() {
            this.eventFn();
        },

        eventFn : function() {
            dom.dbTransferAll.bind('click',function() {
                
                if (this.checked == true) {
                    dom.dbCollections.prop("checked",true);
                } else {
                    dom.dbCollections.prop("checked", false); 
                }
                
            });
        }
    }
    dbTransfer.init();
})();