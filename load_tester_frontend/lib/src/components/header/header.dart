import 'package:angular/angular.dart';
import 'package:angular_components/angular_components.dart';
import 'package:angular_router/angular_router.dart';
import 'package:load_tester_frotend/src/route_paths.dart';
import 'package:load_tester_frotend/src/routes.dart';

@Component(
  selector: 'header',
  directives: const [
    materialInputDirectives,
    MaterialTemporaryDrawerComponent,
    DeferredContentDirective,
    MaterialPersistentDrawerDirective,
    MaterialToggleComponent,
    MaterialButtonComponent,
    MaterialIconComponent,
    routerDirectives,
    NgIf,
  ],
  templateUrl: 'header.html',
  styleUrls: const [
    'header.css',
    'package:angular_components/app_layout/layout.scss.css',
    'package:angular_components/css/mdc_web/card/mdc-card.scss.css',
  ],
  providers: [
    routerProviders,
  ],
  exports: [Routes, RoutePaths],
)
class HeaderComponent {
  bool drawerVisible = false;

  HeaderComponent();

  void navigateCreateRequest() {
    print("Here");
  }
}
