(function() {
    var dom = {
        jsFile : $('#js-file'),
        jsFileList : $('#js-file-list'),
        jsonFile : $('#json-file'),
        jsonFileList : $('#json-file-list')
    }

    var dbImport = {
        init : function() {
            this.eventFn();
        },
        eventFn : function () {
            dom.jsFile.bind('change',function() {
                var fileNames = '';
                $.each(this.files, function() {
                  fileNames += '<span class="am-badge">' + this.name + '</span> ';
                });
                dom.jsFileList.html(fileNames);
            }),

            dom.jsonFile.bind('change',function() {
                var fileNames = '';
                $.each(this.files, function() {
                  fileNames += '<span class="am-badge">' + this.name + '</span> ';
                });
                dom.jsonFileList.html(fileNames);
            })
        }
    }

    dbImport.init();
})()