/** @license React v16.6.0-alpha.8af6728
 * create-subscription.development.js
 *
 * Copyright (c) Facebook, Inc. and its affiliates.
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

'use strict';



if (process.env.NODE_ENV !== "production") {
  (function() {
'use strict';

Object.defineProperty(exports, '__esModule', { value: true });

var React = require('react');

/**
 * Use invariant() to assert state which your program assumes to be true.
 *
 * Provide sprintf-style format (only %s is supported) and arguments
 * to provide information about what broke and what you were
 * expecting.
 *
 * The invariant message will be stripped in production, but the invariant
 * will remain to ensure logic does not differ in production.
 */

var validateFormat = function () {};

{
  validateFormat = function (format) {
    if (format === undefined) {
      throw new Error('invariant requires an error message argument');
    }
  };
}

function invariant(condition, format, a, b, c, d, e, f) {
  validateFormat(format);

  if (!condition) {
    var error = void 0;
    if (format === undefined) {
      error = new Error('Minified exception occurred; use the non-minified dev environment ' + 'for the full error message and additional helpful warnings.');
    } else {
      var args = [a, b, c, d, e, f];
      var argIndex = 0;
      error = new Error(format.replace(/%s/g, function () {
        return args[argIndex++];
      }));
      error.name = 'Invariant Violation';
    }

    error.framesToPop = 1; // we don't care about invariant's own frame
    throw error;
  }
}

// Relying on the `invariant()` implementation lets us
// preserve the format and params in the www builds.

/**
 * Similar to invariant but only logs a warning if the condition is not met.
 * This can be used to log issues in development environments in critical
 * paths. Removing the logging code for production environments will keep the
 * same logic and follow the same code paths.
 */

var warningWithoutStack = function () {};

{
  warningWithoutStack = function (condition, format) {
    for (var _len = arguments.length, args = Array(_len > 2 ? _len - 2 : 0), _key = 2; _key < _len; _key++) {
      args[_key - 2] = arguments[_key];
    }

    if (format === undefined) {
      throw new Error('`warningWithoutStack(condition, format, ...args)` requires a warning ' + 'message argument');
    }
    if (args.length > 8) {
      // Check before the condition to catch violations early.
      throw new Error('warningWithoutStack() currently supports at most 8 arguments.');
    }
    if (condition) {
      return;
    }
    if (typeof console !== 'undefined') {
      var _args$map = args.map(function (item) {
        return '' + item;
      }),
          a = _args$map[0],
          b = _args$map[1],
          c = _args$map[2],
          d = _args$map[3],
          e = _args$map[4],
          f = _args$map[5],
          g = _args$map[6],
          h = _args$map[7];

      var message = 'Warning: ' + format;

      // We intentionally don't use spread (or .apply) because it breaks IE9:
      // https://github.com/facebook/react/issues/13610
      switch (args.length) {
        case 0:
          console.error(message);
          break;
        case 1:
          console.error(message, a);
          break;
        case 2:
          console.error(message, a, b);
          break;
        case 3:
          console.error(message, a, b, c);
          break;
        case 4:
          console.error(message, a, b, c, d);
          break;
        case 5:
          console.error(message, a, b, c, d, e);
          break;
        case 6:
          console.error(message, a, b, c, d, e, f);
          break;
        case 7:
          console.error(message, a, b, c, d, e, f, g);
          break;
        case 8:
          console.error(message, a, b, c, d, e, f, g, h);
          break;
        default:
          throw new Error('warningWithoutStack() currently supports at most 8 arguments.');
      }
    }
    try {
      // --- Welcome to debugging React ---
      // This error was thrown as a convenience so that you can use this stack
      // to find the callsite that caused this warning to fire.
      var argIndex = 0;
      var _message = 'Warning: ' + format.replace(/%s/g, function () {
        return args[argIndex++];
      });
      throw new Error(_message);
    } catch (x) {}
  };
}

var warningWithoutStack$1 = warningWithoutStack;

function _classCallCheck(instance, Constructor) { if (!(instance instanceof Constructor)) { throw new TypeError("Cannot call a class as a function"); } }

function _possibleConstructorReturn(self, call) { if (!self) { throw new ReferenceError("this hasn't been initialised - super() hasn't been called"); } return call && (typeof call === "object" || typeof call === "function") ? call : self; }

function _inherits(subClass, superClass) { if (typeof superClass !== "function" && superClass !== null) { throw new TypeError("Super expression must either be null or a function, not " + typeof superClass); } subClass.prototype = Object.create(superClass && superClass.prototype, { constructor: { value: subClass, enumerable: false, writable: true, configurable: true } }); if (superClass) Object.setPrototypeOf ? Object.setPrototypeOf(subClass, superClass) : subClass.__proto__ = superClass; }

function createSubscription(config) {
  var getCurrentValue = config.getCurrentValue,
      _subscribe = config.subscribe;


  !(typeof getCurrentValue === 'function') ? warningWithoutStack$1(false, 'Subscription must specify a getCurrentValue function') : void 0;
  !(typeof _subscribe === 'function') ? warningWithoutStack$1(false, 'Subscription must specify a subscribe function') : void 0;

  // Reference: https://gist.github.com/bvaughn/d569177d70b50b58bff69c3c4a5353f3
  var Subscription = function (_React$Component) {
    _inherits(Subscription, _React$Component);

    function Subscription() {
      var _temp, _this, _ret;

      _classCallCheck(this, Subscription);

      for (var _len = arguments.length, args = Array(_len), _key = 0; _key < _len; _key++) {
        args[_key] = arguments[_key];
      }

      return _ret = (_temp = (_this = _possibleConstructorReturn(this, _React$Component.call.apply(_React$Component, [this].concat(args))), _this), _this.state = {
        source: _this.props.source,
        value: _this.props.source != null ? getCurrentValue(_this.props.source) : undefined
      }, _this._hasUnmounted = false, _this._unsubscribe = null, _temp), _possibleConstructorReturn(_this, _ret);
    }

    Subscription.getDerivedStateFromProps = function getDerivedStateFromProps(nextProps, prevState) {
      if (nextProps.source !== prevState.source) {
        return {
          source: nextProps.source,
          value: nextProps.source != null ? getCurrentValue(nextProps.source) : undefined
        };
      }

      return null;
    };

    Subscription.prototype.componentDidMount = function componentDidMount() {
      this.subscribe();
    };

    Subscription.prototype.componentDidUpdate = function componentDidUpdate(prevProps, prevState) {
      if (this.state.source !== prevState.source) {
        this.unsubscribe();
        this.subscribe();
      }
    };

    Subscription.prototype.componentWillUnmount = function componentWillUnmount() {
      this.unsubscribe();

      // Track mounted to avoid calling setState after unmounting
      // For source like Promises that can't be unsubscribed from.
      this._hasUnmounted = true;
    };

    Subscription.prototype.render = function render() {
      return this.props.children(this.state.value);
    };

    Subscription.prototype.subscribe = function subscribe() {
      var _this2 = this;

      var source = this.state.source;

      if (source != null) {
        var _callback = function (value) {
          if (_this2._hasUnmounted) {
            return;
          }

          _this2.setState(function (state) {
            // If the value is the same, skip the unnecessary state update.
            if (value === state.value) {
              return null;
            }

            // If this event belongs to an old or uncommitted data source, ignore it.
            if (source !== state.source) {
              return null;
            }

            return { value: value };
          });
        };

        // Store the unsubscribe method for later (in case the subscribable prop changes).
        var unsubscribe = _subscribe(source, _callback);
        !(typeof unsubscribe === 'function') ? invariant(false, 'A subscription must return an unsubscribe function.') : void 0;

        // It's safe to store unsubscribe on the instance because
        // We only read or write that property during the "commit" phase.
        this._unsubscribe = unsubscribe;

        // External values could change between render and mount,
        // In some cases it may be important to handle this case.
        var _value = getCurrentValue(this.props.source);
        if (_value !== this.state.value) {
          this.setState({ value: _value });
        }
      }
    };

    Subscription.prototype.unsubscribe = function unsubscribe() {
      if (typeof this._unsubscribe === 'function') {
        this._unsubscribe();
      }
      this._unsubscribe = null;
    };

    return Subscription;
  }(React.Component);

  return Subscription;
}

exports.createSubscription = createSubscription;
  })();
}
