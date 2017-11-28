package contracts

type App interface {
  Process(request Request) Response
}

type Request interface {
  // getData()
  // getFormat() TransportFormat
}

type Response interface {
  // setData(TransportData)
  // getData() TransportData
  // setFormat(TransportFormat)
  // getFormat() TransportFormat
}

type TransportFormat interface {

}

type TransportData map[string]TransportData


type Initializer interface {
  Prepare(app App)
  Run(app App)
}

type Middleware interface {
  BeforeRequest(app App, request Request)
  AfterRequest(app App, request Request)
}
