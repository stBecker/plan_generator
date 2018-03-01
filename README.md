# plan_generator

# Setup
- Clone repo

- run unit tests:

        go test

- build the binary:

        go build

- run the binary


# Usage
POST a JSON payload to http://localhost:8080/generate-plan

returns a JSON response


Example payload:

    {
    "loanAmount": "5000",
    "nominalRate": "5.0",
    "duration": 24,
    "startDate": "2018-01-01T00:00:01Z"
    }


Example response:

    {
    [
    {
    "borrowerPaymentAmount": "219.36",
    "date": "2018-01-01T00:00:00Z",
    "initialOutstandingPrincipal": "5000.00",
    "interest": "20.83",
    "principal": "198.53",
    "remainingOutstandingPrincipal": "4801.47",
    },
    {
    "borrowerPaymentAmount": "219.36",
    "date": "2018-02-01T00:00:00Z",
    "initialOutstandingPrincipal": "4801.47",
    "interest": "20.00",
    "principal": "199.36",
    "remainingOutstandingPrincipal": "4638",
    },
    ...

