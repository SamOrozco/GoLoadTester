import 'package:angular/angular.dart';
import 'dart:convert';
import 'package:angular_components/angular_components.dart';
import 'package:angular_router/angular_router.dart';
import 'package:load_tester_frotend/src/blocs/request_bloc.dart';
import 'package:load_tester_frotend/src/components/request_url/request_url.dart'
    as url;
import 'package:load_tester_frotend/src/components/components.dart';
import 'package:load_tester_frotend/src/components/configure_schedule/configure_schedule.dart';
import 'package:load_tester_frotend/src/models/models.dart';
import 'package:load_tester_frotend/src/services/services.dart';
import 'package:load_tester_frotend/src/route_paths.dart';

@Component(
  selector: 'add-request',
  directives: const [
    MaterialButtonComponent,
    RequestUrlComponent,
    AddStringPairsComponent,
    ConfigureScheduleComponent,
    MaterialInputComponent,
    MaterialSpinnerComponent,
    HeaderComponent,
    NgIf,
  ],
  templateUrl: 'add_request.html',
  styleUrls: const [
    'add_request.css',
  ],
)
class AddRequest {
  url.RequestUrlComponent requestUrlComponent;
  AddStringPairsComponent headerComponent;
  AddStringPairsComponent queryParams;
  ConfigureScheduleComponent configureScheduleComponent;
  final RequestBloc bloc;
  final Router _router;
  bool loading = false;

  AddRequest(RequestService requestService, this._router)
      : bloc = new RequestBloc(requestService);

  @ViewChild('request_url')
  void set RequestUrlComponent(url.RequestUrlComponent comp) {
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

  @ViewChild('configure_schedule')
  void set configureSchedule(ConfigureScheduleComponent comp) {
    this.configureScheduleComponent = comp;
  }

  Future submitRequest() async {
    this.loading = true;
    var response = await this.bloc.createScheduleRequest(getRequestFromForm());
    print(response.scheduleId);
    // navigate to ID!!!
    var res = await this._router.navigateByUrl(RoutePaths.scheduleDetail.toUrl(
          parameters: {"scheduleId": response.scheduleId},
        ));
    if (res == NavigationResult.BLOCKED_BY_GUARD) {
      print("Blocked by guard");
    }
    if (res == NavigationResult.INVALID_ROUTE) {
      print("Invalid Route");
    }
    if (res == NavigationResult.SUCCESS) {
//      print("Success");
    }
    print(res.toString());
    this.loading = false;
  }

  ScheduleRequest getRequestFromForm() {
    var url = this.requestUrlComponent?.url;
    var requestType = this.requestUrlComponent?.selectedRequestType;
    var headers = this.headerComponent?.keyValuePairs;
    var queryParams = this.queryParams?.keyValuePairs;
    var requestCount = this.configureScheduleComponent.requestCount;
    var timeInterval = this.configureScheduleComponent.timeInterval;
    var timeType = this.configureScheduleComponent.selectedTimeTypeString;
    return new ScheduleRequest(
        url: url,
        requestType: requestType,
        headers: KeyValuePair.getMapFromKeyValues(headers),
        queryParams: KeyValuePair.getMapFromKeyValues(queryParams),
        requestCount: requestCount,
        intervalCount: timeInterval,
        intervalType: timeType);
  }
}
