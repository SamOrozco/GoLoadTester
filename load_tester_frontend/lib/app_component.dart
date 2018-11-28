import 'package:angular/angular.dart';
import 'package:angular_components/angular_components.dart';
import 'package:angular_router/angular_router.dart';
import 'package:load_tester_frotend/src/components/add_request/add_request.dart';
import 'package:load_tester_frotend/src/components/components.dart';
import 'package:load_tester_frotend/src/routes.dart';

// AngularDart info: https://webdev.dartlang.org/angular
// Components info: https://webdev.dartlang.org/components

@Component(
  selector: 'my-app',
  styleUrls: [
    'app_component.css',
    'package:angular_components/css/mdc_web/card/mdc-card.scss.css',
  ],
  templateUrl: 'app_component.html',
  directives: [
    AddRequest,
    materialInputDirectives,
    MaterialTemporaryDrawerComponent,
    DeferredContentDirective,
    MaterialButtonComponent,
    MaterialIconComponent,
    MaterialPersistentDrawerDirective,
    MaterialToggleComponent,
    MaterialListComponent,
    MaterialListItemComponent,
    MaterialButtonComponent,
    MaterialIconComponent,
    NgIf,
    routerDirectives,
    HeaderComponent,
  ],
  providers: [
    materialProviders,
  ],
  exports: [Routes, RoutePaths],
)
class AppComponent {
  final Router _router;
  int pageNum;
  bool drawerVisible;

  AppComponent(this._router) {
    pageNum = 0;
    drawerVisible = false;
  }
// Nothing here yet. All logic is in TodoListComponent.

}
