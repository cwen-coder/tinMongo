(function() {
    var dom = {
        usersListA : $('#users-list-a'),
        usersList : $('#users-list'),
        addUserA : $('#add-user-a'),
        addUser : $('#add-user')
    }

    var dbUsers = {
        init : function() {
            this.eventFn();
        },
        eventFn : function() {
            dom.usersListA.bind('click', function(e){
                e.preventDefault();  
                dom.usersList.show();
                dom.addUser.hide();
            }),

            dom.addUserA.bind('click', function(e) {
                e.preventDefault();
                dom.usersList.hide();
                dom.addUser.show();
            })
        }
    }
    dbUsers.init();
})()