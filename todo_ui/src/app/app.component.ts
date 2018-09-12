import { Component } from '@angular/core';

import { Platform } from '@ionic/angular';
import { SplashScreen } from '@ionic-native/splash-screen/ngx';
import { StatusBar } from '@ionic-native/status-bar/ngx';
import * as A from './../../../todo_ui/api/todo/v2/TodoServiceClientPb';
import {CreateRequest, CreateResponse} from "../../../todo_ui/api/todo/v2/todo_pb";

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

        const client = new A.TodoServiceClient("127.0.0.1:2345", {}, {});
        request = new CreateRequest();
        request.setItem({

        });
        client.create(request, null, (err: grpcWeb.Error, response: CreateResponse) => void {

        })
    });
  }
}
