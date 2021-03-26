<template>
  <div id="email_wrapper">
    <!-- Close Button -->
    <div id="close_btn" @click="closeEmailView">X</div>
    <!-- Content -->
    <div id="email">
      <form @submit.prevent="sendEmail">
        <!-- Email -->
        <section>
          <label>Email</label><br />
          <input class="input_fields" id="other_email" name="email"
            placeholder="Your Email" type="email" v-model="emailAddr" required />
        </section>
        <!-- Message -->
        <section>
          <label>Message</label><br />
          <textarea class="input_fields" id="message" name="email"
            placeholder="Your Message" rows="10" v-model="messageDetail"></textarea>
        </section>
        <!-- Send -->
        <section>
          <input id="submit_btn" type="submit" value="Send" />
        </section>
      </form>
    </div>

    <!-- Loading -->
    <div id="loading" ref="loadingRef" style="display: none;">
      <div id="icon">
        <img :src="Loading" />
      </div>
      <div id="text">
        <h2>Please Wait!!!</h2>
      </div>
    </div>
  </div>
</template>

<script>
import { ref, watch } from 'vue';

// EmailJS
import { init, sendForm } from 'emailjs-com';

// Plugin
import { useStore } from '@/components/Utils/plugin';

// GIF
import Loading from '@/assets/gifs/loading.gif';

export default {
  name: 'Email',
  setup() {
    // Emitter
    const emitter = ref(useStore('emitter'));

    // Bind Variables to The Input Fields
    const emailAddr = ref(null);
    const messageDetail = ref(null);

    const serviceID = ref('service_hdvonwp');
    const templateID = ref('template_qjzmd4e');
    const userID = ref('user_4Bjch1pMfzrpotdZmjxjM');

    // Ref
    const loadingRef = ref(null);

    // Close The View
    function closeEmailView() {
      if (!emitter.value.err) {
        emitter.value.store.emit('email_screen', '');
      }
    }

    // Send Email
    function sendEmail(e) {
      try {
        loadingRef.value.style.display = 'block';
        init(userID.value);
        sendForm(serviceID.value, templateID.value, e.target, userID.value, {
          email: emailAddr.value,
          message: messageDetail.value,
        })
          .then((result) => {
            if (result.status === 200) {
              emailAddr.value = '';
              messageDetail.value = '';
              emitter.value.store.emit('email_screen', '');
              loadingRef.value.style.display = 'none';
            }
          }, (error) => {
            emitter.value.store.emit('email_screen', error);
          });
      } catch (err) {
        emitter.value.store.emit('email_screen', err);
      }
    }

    // Watch CHanges from The Input Variables
    watch(([emailAddr, messageDetail]), ([addr, mes]) => {
      emailAddr.value = addr;
      messageDetail.value = mes;
    });

    return {
      emailAddr,
      messageDetail,
      closeEmailView,
      sendEmail,
      Loading,
      loadingRef,
    };
  },
};
</script>

<style lang="scss" scoped>
#email_wrapper {
  display: block;
  position: relative;
  top: 15%;
  width: 80%;
  padding: 20px;
  text-align: right;

  #close_btn {
    display: inline-block;
    position: relative;
    top: 15%;
    padding: 20px;
    border-radius: 50%;
    border: 4px dashed black;
    cursor: pointer;
    transition: .2s all;

    &:hover {
      border: 4px dashed white;
      color: white;
    }
  }

  #email {
    position: relative;
    margin: 10px;
    padding: 20px;
    color: snow;
    border-radius: 5px;
    background-color: rgba(87, 87, 87, 0.233);
    text-align: left;

    section {
      margin: 20px;
      text-align: center;

      label {
        float: left;
        margin-left: 5px;
      }

      .input_fields {
        margin: 5px;
        width: 100%;
      }

      #submit_btn {
        padding: 15px 40px;
        color: white;
        border: none;
        border-radius: 4px;
        font-size: 1rem;
        background-color: rgb(204, 46, 46);
        cursor: pointer;
        transition: .2s background-color;

        &:hover {
          background-color: rgb(238, 90, 90);
        }
      }
    }
  }

  #loading {
    position: absolute;
    top: 50%;
    left: 50%;
    width: 50%;
    transform: translate(-50%, -50%);
    text-align: center;
    border-radius: 5px;
    background-color: rgba(0, 0, 0, 0.25);

    #icon {
      position: relative;

      img {
        width: 15%;
      }
    }
  }
}
</style>
