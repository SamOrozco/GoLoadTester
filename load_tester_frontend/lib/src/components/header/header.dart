import 'package:angular/angular.dart';
import 'package:angular_components/angular_components.dart';
import 'package:angular_router/angular_router.dart';

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
)
class HeaderComponent {
  bool drawerVisible = false;
  HeaderComponent();
}
