/**
 * Spyderbat API UI & Public APIs
 * Restful APIs for use by UI & customers.
 *
 * The version of the OpenAPI document: 1.0.0
 * Contact: apisupport@spyderbat.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 *
 */

(function(root, factory) {
  if (typeof define === 'function' && define.amd) {
    // AMD.
    define(['expect.js', process.cwd()+'/src/index'], factory);
  } else if (typeof module === 'object' && module.exports) {
    // CommonJS-like environments that support module.exports, like Node.
    factory(require('expect.js'), require(process.cwd()+'/src/index'));
  } else {
    // Browser globals (root is window)
    factory(root.expect, root.SpyderbatApi);
  }
}(this, function(expect, SpyderbatApi) {
  'use strict';

  var instance;

  beforeEach(function() {
    instance = new SpyderbatApi.ApiSOARListHandlerOutput();
  });

  var getProperty = function(object, getter, property) {
    // Use getter method if present; otherwise, get the property directly.
    if (typeof object[getter] === 'function')
      return object[getter]();
    else
      return object[property];
  }

  var setProperty = function(object, setter, property, value) {
    // Use setter method if present; otherwise, set the property directly.
    if (typeof object[setter] === 'function')
      object[setter](value);
    else
      object[property] = value;
  }

  describe('ApiSOARListHandlerOutput', function() {
    it('should create an instance of ApiSOARListHandlerOutput', function() {
      // uncomment below and update the code to test ApiSOARListHandlerOutput
      //var instance = new SpyderbatApi.ApiSOARListHandlerOutput();
      //expect(instance).to.be.a(SpyderbatApi.ApiSOARListHandlerOutput);
    });

    it('should have the property investigateSourceUrl (base name: "investigate_source_url")', function() {
      // uncomment below and update the code to test the property investigateSourceUrl
      //var instance = new SpyderbatApi.ApiSOARListHandlerOutput();
      //expect(instance).to.be();
    });

    it('should have the property source (base name: "source")', function() {
      // uncomment below and update the code to test the property source
      //var instance = new SpyderbatApi.ApiSOARListHandlerOutput();
      //expect(instance).to.be();
    });

  });

}));
