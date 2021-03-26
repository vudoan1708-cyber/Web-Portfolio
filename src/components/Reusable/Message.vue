<template>
  <div id="message_wrapper">
    <div id="message">
      <!-- Error Message -->
      <section>
        <div id="error">
          <h2>Error</h2>
        </div>
      </section>
      <!-- Message -->
      <section>
        <div id="message_detail">
          <p>{{ details }}</p>
        </div>
      </section>
      <!-- Proceed Button -->
      <div id="proceed" @click="closeErrorBoard">
        <p>Got It</p>
      </div>
    </div>
  </div>
</template>

<script>
import { ref } from 'vue';

// Plugin
import { useStore } from '@/components/Utils/plugin';

export default {
  name: 'Message',
  props: {
    emailDetail: {
      type: String,
    },
  },
  setup(props) {
    // Emitter
    const emitter = ref(useStore('emitter'));

    const details = ref(props.emailDetail);

    // Close The Erro Board
    function closeErrorBoard() {
      if (!emitter.value.err) {
        emitter.value.store.emit('email_screen', details.value);
      }
    }

    return {
      details,
      closeErrorBoard,
    };
  },
};
</script>

<style lang="scss" scoped>
#message_wrapper {
  position: absolute;
  top: 0;
  width: 80%;
  height: 100%;
  margin: 0;
  z-index: 5;
  background-color: rgba(0, 0, 0, .5);

  #message {
    position: absolute;
    top: 50%;
    left: 50%;
    transform: translate(-50%, -50%);
    width: 50%;
    height: 25%;
    background-color: rgb(235, 34, 34);
    border-radius: 10px;
    box-shadow: 0px 0px 10px snow;
    color: snow;
    text-align: center;

    section {
      margin: 10px;
    }

    #proceed {
      position: absolute;
      bottom: 0;
      left: 50%;
      transform: translateX(-50%);
      padding: 20px;
      border-radius: 5px;
      background-color: rgb(73, 73, 73);
      cursor: pointer;

      &:hover {
        background-color: rgb(71, 194, 71);
      }
    }
  }
}
</style>
