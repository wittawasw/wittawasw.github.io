---
layout: post
title: รูปแบบ comment ที่เป็นประโยชน์ในการเขียนโปรแกรมร่วมกัน
date: 2022-02-20 17:31 +0700
---
การ comment คือ การทำให้ประโยคหรือตัวอักษรที่เราเขียนในโปรแกรมไม่ถูกประมวลผล โดยอาจมีไว้เพื่อให้บอท
(bot) หรือโปรแกรมเมอร์คนอื่นอ่าน เป็นรูปแบบการการสื่อสารแบบหนึ่งที่โปรแกรมเมอร์สามารถสื่อสารได้
เพื่อให้การทำงานร่วมกันมีประสิทธิภาพมากขึ้น

> ไม่จำเป็นต้อง comment ก็ได้ ถ้าดีวางรูปแบบและชื่อตัวแปรของ code ดีพอ

คำพูดนี้เป็นประโยคที่พบได้บ่อยในการแลกเปลี่ยนกันในหัวข้อ **เราควร comment ใน code หรือไม่?**
ในความเห็นของผมคือ การ comment ยังมีความจำเป็นในบางโอกาสและมีประโยชน์มากกว่าโทษ
ในบางโอกาสที่ถึงแม้จะไม่มีความจำเป็นแต่ก็จะเป็นประโยชน์กับเพื่อนร่วมงานมากกว่าการที่เราไม่ comment
เพราะคำอธิบายที่เป็นคำพูด ส่วนมากแล้วจะเข้าใจได้ง่ายกว่า

ในการทำงานร่วมกันกับคนหลายคน ถึงแม้ว่าผู้ที่เขียน code นั้น
จะเชื่อว่าสิ่งที่เขียนลงไป เป็นสิ่งที่ง่ายสำหรับผู้เขียน แต่ทุกคนมีที่มาและพื้นฐานของความรู้ที่แตกต่างกัน
หากเราหลีกเลี่ยงการ comment ในส่วนที่สำคัญ นอกจากจะไม่เกิดประโยชน์ ยังอาจทำให้การทำงานร่วมกัน
สามารถทำได้ช้าลงเนื่องจากเพื่อนร่วมงานอาจจะไม่สามารถทำความเข้าใจกับ code ได้

โดยปกติของการ comment ไม่ได้มีการระบุหลักการหรือรูปแบบที่ตายตัวเอาไว้ ทำให้ผู้ที่เริ่มต้นเขียนโปรแกรม
มักจะไม่กล้าเริ่มเขียน comment เพราะกลัวว่าจะทำให้ code ของตนเองดูไม่ดี แต่หลักคิดของผมคือ
หากโปรแกรมเมอร์ที่มาเขียนโปรแกรมต่อจากเรา จำเป็นต้องได้รับคำอธิบายก่อนการเริ่มทำงาน
ไม่ว่าจะเรื่องอะไร คำอธิบายเรื่องนั้นก็สมควรถูกใส่เอาไว้ใน comment ของ code เช่นตัวอย่างต่อไปนี้

### อธิบายความสำคัญหรือหน้าที่ของ class, module หรือ namespace

ในการเขียนโปรแกรมที่มีการแบ่งส่วนการทำงานออกเป็น class, module หรือ namespace
หลายๆครั้งจะตั้งชื่อด้วยคำที่มีความหมายทั่วไป เพื่อให้สามารถคาดเดาการทำงานได้จากชื่อ
แต่บางครั้งชื่อที่ใช้ก็อาจจะมีควาหมายที่กว้าง และสามารถใช้ได้กับหลายอย่าง

ยกตัวอย่าง comment จาก source code ของ `Redis` ในส่วน `Connection` ที่หากเปิดไฟล์ขึ้นมาอ่านครั้งแรก
อาจเข้าใจว่า `Connection` ในที่นี้หมายถึงการตั้งค่า network connection เช่น การระบุขนาด buffer ,
การระบุชื่อ socket ที่ใช้  ผู้เขียนจึงได้ comment เอาไว้ว่าไฟล์ `connection.c` นี้เป็นเพียง abstraction
ให้เรียกใช้เท่านั้น ส่วนของการตั้งค่าต่างๆของ network อยู่ใน `networking.c`

{% highlight c %}
#include "server.h"
#include "connhelpers.h"

/* The connections module provides a lean abstraction of network connections
 * to avoid direct socket and async event management across the Redis code base.
 *
 * It does NOT provide advanced connection features commonly found in similar
 * libraries such as complete in/out buffer management, throttling, etc. These
 * functions remain in networking.c.
 *
 * The primary goal is to allow transparent handling of TCP and TLS based
 * connections. To do so, connections have the following properties:
 *
 * 1. A connection may live before its corresponding socket exists.  This
 *    allows various context and configuration setting to be handled before
 *    establishing the actual connection.
 * 2. The caller may register/unregister logical read/write handlers to be
 *    called when the connection has data to read from/can accept writes.
 *    These logical handlers may or may not correspond to actual AE events,
 *    depending on the implementation (for TCP they are; for TLS they aren't).
 */

 ConnectionType CT_Socket;
 .
 .
{% endhighlight %}
ที่มา: [redis/src/connection.c#L33](https://github.com/redis/redis/blob/9b0fd9f4d0e7e41f3cc78b4bd37619b574246aef/src/connection.c#L33)

### อธิบายการทำงานของ code

เราสามารถตีความจุดประสงค์ของผู้เขียนโปรแกรมจากโปรแกรมที่เขียนได้
แต่กระบวนการนั้นต้องใช้ระยะเวลาและความเชี่ยวชาญที่บางคนในทีมไม่มี การเขียนอธิบายการทำงานคร่าวๆเอาไว้
นอกจากจะช่วยประหยัดเวลาของโปรแกรมเมอร์ที่อาจถูกมอบหมายให้ทำงานต่อจากเรา
ยังช่วยลดปัญหาที่อาจเกิดจากการตีความผิดอีกด้วย

ลองอ่านตัวอย่าง code ด้านล่างดูโดยเริ่มจาก comment จะเห็นประโยชน์ชัดเจนว่าการได้อ่าน comment ที่อธิบายการทำงานก่อนทำให้สามารถเข้าใจ code ได้เร็วขึ้น และยังป้องกันความผิดพลาดจากการตีความผิด

{% highlight python %}
def call_with_retry(self, do, fail):
    """
    Execute an operation that might fail and returns its result, or
    raise the exception that was thrown depending on the `Backoff` object.
    `do`: the operation to call. Expects no argument.
    `fail`: the failure handler, expects the last error that was thrown
    """
    self._backoff.reset()
    failures = 0
    while True:
        try:
            return do()
        except self._supported_errors as error:
            failures += 1
            fail(error)
            if failures > self._retries:
                raise error
            backoff = self._backoff.compute(failures)
            if backoff > 0:
                sleep(backoff)
{% endhighlight %}
ที่มา: [redis-py/redis/retry.py#L34](https://github.com/redis/redis-py/blob/e3c989d93e914e6502bd5a72f15ded49a135c5be/redis/retry.py#L34)

> ในกรณีที่เป็น lib ให้มีการเรียกใช้จากภายนอกได้ การยกตัวอย่างใน comment ก็เป็นทางเลือกที่ดี

{% highlight python %}
class AsyncDataLoaderMixin(object):
    """
    Async Mixin on top of implementation of BaseDataLoader. It contains a seperate thread
    which reads batch from self._iterate() and push them in the queue. The self.__iter__() function
    will pop the batch from the queue.
    If async_loader_queue_size is set to 0, the data loader will not work in async mode.
    For example:
        class PytorchAsyncDataLoader(AsyncDataLoaderMixin, PytorchDataLoader):
    """

    def __init__(self, async_loader_queue_size=64, debug_data_loader=False, *args, **kwargs):
        """
        initialize the async data loader. Need to add this in the __init__() of the implementation
        """
        self.async_loader_queue_size = async_loader_queue_size
        self.debug_data_loader = debug_data_loader
        super().__init__(*args, **kwargs)

        print(f"Apply the AsyncDataLoaderMixin on top of the data loader, async_loader_queue_size={async_loader_queue_size}. ")

        if self.async_loader_queue_size > 0:
            self.finished_event = Event()
            self.queue = Queue(self.async_loader_queue_size)
            self.thread = Thread(target=self._async_worker)
            self.thread.daemon = True
            self.started = False
{% endhighlight %}
ที่มา: [horovod/data/data_loader_base.py#L49](https://github.com/horovod/horovod/blob/046c071cb6b0dfbeb6082c6bdf2366e2bf7a6ad9/horovod/data/data_loader_base.py#L49)

### อธิบายความหมายของตัวแปรแต่ล่ะตัว

ชื่อตัวแปรก็อยู่ในแนวคิดที่ใกล้เคียงกับการอธิบายชื่อ module, class โดยมีประโยชน์มากสำหรับผู้เริ่มต้น
หรือเพื่อนร่วมงานที่เพิ่งเริ่มต้นทำงานร่วมกัน ให้สามารถรู้ถึงความสำคัญของตัวแปรแต่ล่ะตัวได้
ว่ามีผลกระทบอะไรบ้างกับโปรแกรม

{% highlight ruby %}
# Create a new classifier with multi-layer preceptron.
#
# @param hidden_units [Array] The number of units in the i-th hidden layer.
# @param dropout_rate [Float] The rate of the units to drop.
# @param learning_rate [Float] The initial value of learning rate in Adam optimizer.
# @param decay1 [Float] The smoothing parameter for the first moment in Adam optimizer.
# @param decay2 [Float] The smoothing parameter for the second moment in Adam optimizer.
# @param max_iter [Integer] The maximum number of epochs that indicates
#   how many times the whole data is given to the training process.
# @param batch_size [Intger] The size of the mini batches.
# @param tol [Float] The tolerance of loss for terminating optimization.
# @param verbose [Boolean] The flag indicating whether to output loss during iteration.
# @param random_seed [Integer] The seed value using to initialize the random generator.
def initialize(hidden_units: [128, 128], dropout_rate: 0.4, learning_rate: 0.001, decay1: 0.9, decay2: 0.999,
                max_iter: 200, batch_size: 50, tol: 1e-4, verbose: false, random_seed: nil)
  check_params_type(Array, hidden_units: hidden_units)
  check_params_numeric(dropout_rate: dropout_rate, learning_rate: learning_rate, decay1: decay1, decay2: decay2,
                        max_iter: max_iter, batch_size: batch_size, tol: tol)
  check_params_boolean(verbose: verbose)
  check_params_numeric_or_nil(random_seed: random_seed)
  super
  @classes = nil
  @network = nil
end
{% endhighlight %}

ที่มา: [yoshoku/rumale/lib/rumale/neural_network/mlp_classifier.rb#L49](https://github.com/yoshoku/rumale/blob/main/lib/rumale/neural_network/mlp_classifier.rb#L49)

### อธิบายที่มาของแนวคิดหรือทฤษฎีบทที่ใช้ในการทำงาน

การเขียนโปรแกรม โดยเฉพาะโปรแกรมที่เขียนโดยอ้างอิงมาจากทฤษฎีต่างๆ การ comment
ระบุทฤษฎีที่เกี่ยวข้องเอาไว้ จะช่วยให้โปรแกรมเมอร์สามารถสืบค้นย้อนกลับไปได้
หากต้องการแก้ไขปัญหาที่อาจจะไม่ได้มาจาก error ของโปรแกรม แต่มาจากความผิดพลาดในการนำทฤษฎีมา implement


{% highlight ruby %}
module Rumale
  module NeuralNetwork
    # MLPClassifier is a class that implements classifier based on multi-layer perceptron.
    # MLPClassifier use ReLu as the activation function and Adam as the optimization method
    # and softmax cross entropy as the loss function.
    #
    # @example
    #   estimator = Rumale::NeuralNetwork::MLPClassifier.new(hidden_units: [100, 100], dropout_rate: 0.3)
    #   estimator.fit(training_samples, traininig_labels)
    #   results = estimator.predict(testing_samples)
    class MLPClassifier < BaseMLP
      include Base::Classifier
      # .
      # .
    end
  end
end
{% endhighlight %}

ที่มา: [yoshoku/rumale/lib/rumale/neural_network/mlp_classifier.rb#L9](https://github.com/yoshoku/rumale/blob/345d02d18c596aa496549ad40cbbda0303dea19b/lib/rumale/neural_network/mlp_classifier.rb#L9)



> หรือถ้าแค่จะบอกว่า copy มาจาก web ก็ไม่จำเป็นต้องอธิบายเยอะ แค่ใส่ URL ไว้ก็พอ

{% highlight javascript %}
// http://beeker.io/jquery-document-ready-equivalent-vanilla-javascript
function documentReady(callback) {
  if (document.readyState === "interactive" || document.readyState === "complete") {
    setTimeout(callback, 0);
  } else {
    document.addEventListener("DOMContentLoaded", callback);
  }
}
{% endhighlight %}

ที่มา: [ankane/ahoy/vendor/assets/javascripts/ahoy.js#L178](https://github.com/ankane/ahoy/blob/1cfff6ae3cd96af87574393bf63f5a294e33dd46/vendor/assets/javascripts/ahoy.js#L178)

