package Emitter

import (
	"github.com/gopherjs/gopherjs/js"
	nodejs "github.com/oskca/gopherjs-nodejs"
)

type Listener func(em *Emitter, args ...*js.Object)

func New(obj ...*js.Object) *Emitter {
	em := new(Emitter)
	if len(obj) > 0 {
		em.Object = obj[0]
	} else {
		em.Object = nodejs.Require("events").New()
	}
	return em
}

type Emitter struct {
	*js.Object
	// emitter.addListener(eventName, listener)#
	// Added in: v0.1.26
	// Alias for emitter.on(eventName, listener).

	// emitter.emit(eventName[, ...args])#
	//
	// Added in: v0.1.26
	// Synchronously calls each of the listeners registered for the event named eventName, in the order they were registered, passing the supplied arguments to each.
	//
	// Returns true if the event had listeners, false otherwise.
	Emit func(eventName string, args ...interface{}) bool `js:"emit"`

	// emitter.eventNames()#
	//
	// Added in: v6.0.0
	// Returns an array listing the events for which the emitter has registered listeners. The values in the array will be strings or Symbols.
	//
	// const Emitter = require('events');
	// const myEE = new Emitter();
	// myEE.on('foo', () => {});
	// myEE.on('bar', () => {});
	//
	// const sym = Symbol('symbol');
	// myEE.on(sym, () => {});
	//
	// console.log(myEE.eventNames());
	// // Prints: [ 'foo', 'bar', Symbol(symbol) ]
	EventNames func() []string `js:"eventNames"`

	// emitter.getMaxListeners()#

	// Added in: v1.0.0
	// Returns the current max listener value for the Emitter which is either set by emitter.setMaxListeners(n) or defaults to Emitter.defaultMaxListeners.
	//
	// emitter.listenerCount(eventName)#
	LinstenerCount func(eventName string) int `js:"listenerCount"`

	// Added in: v3.2.0
	// eventName <String> | <Symbol> The name of the event being listened for
	// Returns the number of listeners listening to the event named eventName.

	// emitter.listeners(eventName)#

	// Added in: v0.1.26
	// Returns a copy of the array of listeners for the event named eventName.

	// server.on('connection', (stream) => {
	//   console.log('someone connected!');
	// });
	// console.log(util.inspect(server.listeners('connection')));
	// // Prints: [ [Function] ]

	// emitter.on(eventName, listener)#
	//
	// Added in: v0.1.101
	// eventName <String> | <Symbol> The name of the event.
	// listener <Function> The callback function
	// Adds the listener function to the end of the listeners array for the event named eventName. No checks are made to see if the listener has already been added. Multiple calls passing the same combination of eventName and listener will result in the listener being added, and called, multiple times.
	//
	// server.on('connection', (stream) => {
	//   console.log('someone connected!');
	// });
	// Returns a reference to the Emitter, so that calls can be chained.
	//
	// By default, event listeners are invoked in the order they are added. The emitter.prependListener() method can be used as an alternative to add the event listener to the beginning of the listeners array.
	//
	// const myEE = new Emitter();
	// myEE.on('foo', () => console.log('a'));
	// myEE.prependListener('foo', () => console.log('b'));
	// myEE.emit('foo');
	// // Prints:
	// //   b
	// //   a
	on func(eventName string, listener interface{}) `js:"on"`

	// emitter.once(eventName, listener)#
	//
	// Added in: v0.3.0
	// eventName <String> | <Symbol> The name of the event.
	// listener <Function> The callback function
	// Adds a one time listener function for the event named eventName. The next time eventName is triggered, this listener is removed and then invoked.
	//
	// server.once('connection', (stream) => {
	//   console.log('Ah, we have our first user!');
	// });
	// Returns a reference to the Emitter, so that calls can be chained.
	//
	// By default, event listeners are invoked in the order they are added. The emitter.prependOnceListener() method can be used as an alternative to add the event listener to the beginning of the listeners array.
	//
	// const myEE = new Emitter();
	// myEE.once('foo', () => console.log('a'));
	// myEE.prependOnceListener('foo', () => console.log('b'));
	// myEE.emit('foo');
	// // Prints:
	// //   b
	// //   a
	once func(eventName string, listener interface{}) `js:"once"`

	// emitter.prependListener(eventName, listener)#

	// Added in: v6.0.0
	// eventName <String> | <Symbol> The name of the event.
	// listener <Function> The callback function
	// Adds the listener function to the beginning of the listeners array for the event named eventName. No checks are made to see if the listener has already been added. Multiple calls passing the same combination of eventName and listener will result in the listener being added, and called, multiple times.

	// server.prependListener('connection', (stream) => {
	//   console.log('someone connected!');
	// });
	// Returns a reference to the Emitter, so that calls can be chained.

	// emitter.prependOnceListener(eventName, listener)#

	// Added in: v6.0.0
	// eventName <String> | <Symbol> The name of the event.
	// listener <Function> The callback function
	// Adds a one time listener function for the event named eventName to the beginning of the listeners array. The next time eventName is triggered, this listener is removed, and then invoked.

	// server.prependOnceListener('connection', (stream) => {
	//   console.log('Ah, we have our first user!');
	// });
	// Returns a reference to the Emitter, so that calls can be chained.

	// emitter.removeAllListeners([eventName])#
	//
	// Added in: v0.1.26
	// Removes all listeners, or those of the specified eventName.
	//
	// Note that it is bad practice to remove listeners added elsewhere in the code, particularly when the Emitter instance was created by some other component or module (e.g. sockets or file streams).
	//
	// Returns a reference to the Emitter, so that calls can be chained.
	RemoveAllListener func(eventName ...string) `js:"removeAllListeners"`

	// emitter.removeListener(eventName, listener)#
	//
	// Added in: v0.1.26
	// Removes the specified listener from the listener array for the event named eventName.
	//
	// var callback = (stream) => {
	//   console.log('someone connected!');
	// };
	// server.on('connection', callback);
	// // ...
	// server.removeListener('connection', callback);
	// removeListener will remove, at most, one instance of a listener from the listener array. If any single listener has been added multiple times to the listener array for the specified eventName, then removeListener must be called multiple times to remove each instance.

	// Note that once an event has been emitted, all listeners attached to it at the time of emitting will be called in order. This implies that any removeListener() or removeAllListeners() calls after emitting and before the last listener finishes execution will not remove them from emit() in progress. Subsequent events will behave as expected.

	// const myEmitter = new MyEmitter();

	// var callbackA = () => {
	//   console.log('A');
	//   myEmitter.removeListener('event', callbackB);
	// };

	// var callbackB = () => {
	//   console.log('B');
	// };

	// myEmitter.on('event', callbackA);

	// myEmitter.on('event', callbackB);

	// // callbackA removes listener callbackB but it will still be called.
	// // Internal listener array at time of emit [callbackA, callbackB]
	// myEmitter.emit('event');
	// // Prints:
	// //   A
	// //   B

	// // callbackB is now removed.
	// // Internal listener array [callbackA]
	// myEmitter.emit('event');
	// // Prints:
	// //   A
	// Because listeners are managed using an internal array, calling this will change the position indices of any listener registered after the listener being removed. This will not impact the order in which listeners are called, but it means that any copies of the listener array as returned by the emitter.listeners() method will need to be recreated.

	// Returns a reference to the Emitter, so that calls can be chained.

	// emitter.setMaxListeners(n)#

	// Added in: v0.3.5
	// By default Emitters will print a warning if more than 10 listeners are added for a particular event. This is a useful default that helps finding memory leaks. Obviously, not all events should be limited to just 10 listeners. The emitter.setMaxListeners() method allows the limit to be modified for this specific Emitter instance. The value can be set to Infinity (or 0) to indicate an unlimited number of listeners.

	// Returns a reference to the Emitter, so that calls can be chained.
	PreventDefault func() `js:"preventDefault"`
}

// OnEvent wraps Emitter.on with *Emitter as the first arguments(this in JS)
func (e *Emitter) OnEvent(eventName string, listener Listener) *Emitter {
	// fn := js.MakeFunc(func(this *js.Object, args []*js.Object) interface{} {
	// 	evt := New(this)
	// 	listener(evt, args...)
	// 	return 0
	// })
	fn := func(args ...*js.Object) {
		listener(e, args...)
	}
	e.on(eventName, fn)
	return e
}

// On is a simplified version of OnEvent, using no this
func (e *Emitter) On(eventName string, listener func(args ...*js.Object)) *Emitter {
	e.on(eventName, listener)
	return e
}

// OnceEvent wraps once with `this` support
func (e *Emitter) OnceEvent(eventName string, listener Listener) *Emitter {
	fn := func(args ...*js.Object) {
		listener(e, args...)
	}
	e.once(eventName, fn)
	return e
}

// Once is a simplified version of OnEvent, using no this
func (e *Emitter) Once(eventName string, listener interface{}) *Emitter {
	e.once(eventName, listener)
	return e
}
