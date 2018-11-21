import 'package:angular/angular.dart';
import 'package:angular_components/angular_components.dart';
import 'package:load_tester_frotend/src/components/add_request/add_request.dart';

// AngularDart info: https://webdev.dartlang.org/angular
// Components info: https://webdev.dartlang.org/components

@Component(
  selector: 'my-app',
  styleUrls: [
    'app_component.css',
    'package:angular_components/app_layout/layout.scss.css'
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
    NgIf
  ],
  providers: [
    materialProviders,
  ],
)
class AppComponent {
  int pageNum;
  bool drawerVisible;

  AppComponent() {
    pageNum = 0;
    drawerVisible = false;
  }
// Nothing here yet. All logic is in TodoListComponent.

}
