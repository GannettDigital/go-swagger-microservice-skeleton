# {NAME}
===================

Describe me!

## Configuration

This project comes with a demo swagger.yaml file.  Before you can
generate the goswagger code, you will need to define your own swagger.yaml.


## Generating goswagger code

To generate the models and service code run the following commands:

```
make
```

This generates all the service boilerplate code.

Follow the instructions at https://goswagger.io/generate/server.html
to implement your service.

## Running the demo server

After

  ```
  make run
  ```

## Testing

To run the unit tests on the packages included as a part of this application, run the following command:

  ```
  make test
  ```

After making changes to any of these packages, please update the testing as needed, and verify by running the above command.
