(function() {
    var dom = {
        dbExportAll : $('#dbExportAll'),
        dbCollections : $("form div:first input[name='collection']")
    }

    var dbTransfer = {
        init : function() {
            this.eventFn();
        },

        eventFn : function() {
            dom.dbExportAll.bind('click',function() {
                
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