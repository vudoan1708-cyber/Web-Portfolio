<template>
  <div id="contact">
    <div id="outer_wrapper">
      <div class="img_container" id="mail">
        <img src="@/assets/mail.png" />
      </div>
      <div class="img_container" id="youtube">
        <img src="@/assets/youtube.png" />
      </div>
      <div class="img_container" id="linkedin">
        <img src="@/assets/linkedin.png" />
      </div>
    </div>

    <Headers />
  </div>
</template>

<script>
import Headers from '@/components/Reusable/Headers.vue';

export default {
  name: 'ContactDetail',
  components: {
    Headers,
  },
  setup() {

  },
};
</script>

<style lang="scss" scoped>
#contact {
  position: fixed;
  top: 0;
  left: 0;
  width: 20%;
  height: 100vh;
  background-color: rgb(226, 91, 91);
  box-shadow: 2px 2px 15px black;
  z-index: 0;

  #outer_wrapper {
    position: absolute;
    top: 50%;
    transform: translateY(-50%);

    .img_container {
      width: 20%;
      margin: 40px;
      cursor: pointer;

      img {
        width: 100%;
        transition: 0.2s filter;
      }

      &:hover {
        img {
          filter: invert(100%);
        }
      }
    }
  }
}
</style>
