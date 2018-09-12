import { Component } from '@angular/core';

import { Platform } from '@ionic/angular';
import { SplashScreen } from '@ionic-native/splash-screen/ngx';
import { StatusBar } from '@ionic-native/status-bar/ngx';
import {TodoServiceClient} from './../../../todo_ui/api/todo/v2/TodoServiceClientPb';
import {CreateRequest, CreateResponse} from '../../../todo_ui/api/todo/v2/todo_pb';
import * as grpcWeb from 'grpc-web';

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

        const client = new TodoServiceClient("127.0.0.1:2345", {}, {});
        let request = new CreateRequest();
        request.setItem({
            "a": 1
        });
        client.create(request, null, (err: grpcWeb.Error, response: CreateResponse) => void {

        })
    });
  }
}
