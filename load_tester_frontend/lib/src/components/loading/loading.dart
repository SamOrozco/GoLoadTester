import 'package:angular/angular.dart';
import 'package:angular_components/angular_components.dart';

@Component(
  selector: 'loading',
  directives: const [
    MaterialSpinnerComponent,
  ],
  templateUrl: 'loading.html',
  styleUrls: const [
    'loading.css',
  ],
)
class LoadingComponent {
  LoadingComponent();
}
