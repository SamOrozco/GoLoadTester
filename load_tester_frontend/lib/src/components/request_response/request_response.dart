import 'package:angular/angular.dart';
import 'package:angular_components/angular_components.dart';
import 'package:load_tester_frotend/src/components/components.dart';
import 'package:load_tester_frotend/src/models/models.dart';

@Component(
  selector: 'request-response',
  directives: const [
    MaterialInputComponent,
    ShowCodeComponent,
  ],
  templateUrl: 'request_response.html',
  styleUrls: const [
    'request_response.css',
  ],
)
class RequestResponseComponent {
  @Input()
  RequestResponse request;

  RequestResponseComponent();
}
