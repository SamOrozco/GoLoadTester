import 'package:angular/angular.dart';
import 'package:angular_components/angular_components.dart';
import 'package:load_tester_frotend/src/blocs/request_bloc.dart';
import 'package:load_tester_frotend/src/components/components.dart';
import 'package:load_tester_frotend/src/components/configure_schedule/configure_schedule.dart';

@Component(
  selector: 'add-request',
  directives: const [
    MaterialButtonComponent,
    RequestUrlComponent,
    AddStringPairsComponent,
    ConfigureScheduleComponent,
  ],
  templateUrl: 'add_request.html',
  styleUrls: const [
    'add_request.css',
  ],
)
class AddRequest {
  RequestUrlComponent requestUrlComponent;
  AddStringPairsComponent headerComponent;
  AddStringPairsComponent queryParams;
  final RequestBloc bloc;

  AddRequest(this.bloc);

  @ViewChild('request_url')
  void set RequestUrlComponent(RequestUrlComponent comp) {
    this.requestUrlComponent = comp;
  }

  @ViewChild('add_headers')
  void set HeadComponent(AddStringPairsComponent comp) {
    this.headerComponent = comp;
    this.headerComponent.name = "Add Headers";
  }

  @ViewChild('add_query_params')
  void set QueryParamComponent(AddStringPairsComponent comp) {
    this.queryParams = comp;
    this.queryParams.name = "Add Query Parameters";
  }
}
