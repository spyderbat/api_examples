/**
 * Spyderbat API UI & Public APIs
 * Restful APIs for use by UI & customers.
 *
 * The version of the OpenAPI document: 1.0.0
 * 
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
    factory(root.expect, root.Sbapi);
  }
}(this, function(expect, Sbapi) {
  'use strict';

  var instance;

  beforeEach(function() {
    instance = new Sbapi.DaoAgentLog();
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

  describe('DaoAgentLog', function() {
    it('should create an instance of DaoAgentLog', function() {
      // uncomment below and update the code to test DaoAgentLog
      //var instance = new Sbapi.DaoAgentLog();
      //expect(instance).to.be.a(Sbapi.DaoAgentLog);
    });

    it('should have the property agentRegistrationUid (base name: "agent_registration_uid")', function() {
      // uncomment below and update the code to test the property agentRegistrationUid
      //var instance = new Sbapi.DaoAgentLog();
      //expect(instance).to.be();
    });

    it('should have the property agentUid (base name: "agent_uid")', function() {
      // uncomment below and update the code to test the property agentUid
      //var instance = new Sbapi.DaoAgentLog();
      //expect(instance).to.be();
    });

    it('should have the property bytesSent (base name: "bytes_sent")', function() {
      // uncomment below and update the code to test the property bytesSent
      //var instance = new Sbapi.DaoAgentLog();
      //expect(instance).to.be();
    });

    it('should have the property errUid (base name: "err_uid")', function() {
      // uncomment below and update the code to test the property errUid
      //var instance = new Sbapi.DaoAgentLog();
      //expect(instance).to.be();
    });

    it('should have the property error (base name: "error")', function() {
      // uncomment below and update the code to test the property error
      //var instance = new Sbapi.DaoAgentLog();
      //expect(instance).to.be();
    });

    it('should have the property ipAddress (base name: "ip_address")', function() {
      // uncomment below and update the code to test the property ipAddress
      //var instance = new Sbapi.DaoAgentLog();
      //expect(instance).to.be();
    });

    it('should have the property msg (base name: "msg")', function() {
      // uncomment below and update the code to test the property msg
      //var instance = new Sbapi.DaoAgentLog();
      //expect(instance).to.be();
    });

    it('should have the property orgUid (base name: "org_uid")', function() {
      // uncomment below and update the code to test the property orgUid
      //var instance = new Sbapi.DaoAgentLog();
      //expect(instance).to.be();
    });

    it('should have the property runtimeDetails (base name: "runtime_details")', function() {
      // uncomment below and update the code to test the property runtimeDetails
      //var instance = new Sbapi.DaoAgentLog();
      //expect(instance).to.be();
    });

    it('should have the property sourceUid (base name: "source_uid")', function() {
      // uncomment below and update the code to test the property sourceUid
      //var instance = new Sbapi.DaoAgentLog();
      //expect(instance).to.be();
    });

    it('should have the property time (base name: "time")', function() {
      // uncomment below and update the code to test the property time
      //var instance = new Sbapi.DaoAgentLog();
      //expect(instance).to.be();
    });

    it('should have the property type (base name: "type")', function() {
      // uncomment below and update the code to test the property type
      //var instance = new Sbapi.DaoAgentLog();
      //expect(instance).to.be();
    });

  });

}));
