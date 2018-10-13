import {Component} from '@angular/core';

import {Platform} from '@ionic/angular';
import {SplashScreen} from '@ionic-native/splash-screen/ngx';
import {StatusBar} from '@ionic-native/status-bar/ngx';


@Component({
    selector: 'app-root',
    templateUrl: 'app.component.html'
})
export class AppComponent {
    constructor(
        private platform: Platform,
        private splashScreen: SplashScreen,
        private statusBar: StatusBar
    ) {
        this.initializeApp();
    }

    initializeApp() {
        this.platform.ready().then(() => {
            this.statusBar.styleDefault();
            this.splashScreen.hide();

            // let request = new CreateRequest();
            // // let todo = new Todo();
            // // todo.setTitle("hi");
            // request.setItem({
            //     "a": 1,
            // });
            // const service = new TodoServiceClient("127.0.0.1:2345", {}, {});
            // service.create(request, {}, (err: grpcWeb.Error, response: CreateResponse) => void {
            //
            // });

        });
    }
}
