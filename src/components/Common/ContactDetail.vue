<template>
  <div id="contact">
    <div id="outer_wrapper">
      <div class="img_container" id="mail" @click="openNewTab(Contact.email)">
        <img src="@/assets/logos/mail.png" />
      </div>
      <div class="img_container" id="youtube" @click="openNewTab(Contact.youtube)">
        <img src="@/assets/logos/youtube.png" />
      </div>
      <div class="img_container" id="linkedin" @click="openNewTab(Contact.linkedin)">
        <img src="@/assets/logos/linkedin.png" />
      </div>
      <div class="img_container" id="github" @click="openNewTab(Contact.github)">
        <img src="@/assets/logos/github.png" />
      </div>
    </div>

    <Headers />
    <Footer />
  </div>
</template>

<script>
// Common
import Footer from '@/components/Common/Footer.vue';

// Reusable
import Headers from '@/components/Reusable/Headers.vue';

// JSON
import Contact from '@/components/JSON/contact.json';

export default {
  name: 'ContactDetail',
  components: {
    Headers,
    Footer,
  },
  setup() {
    function openNewTab(tab) {
      window.open(tab);
    }
    return {
      Contact,
      openNewTab,
    };
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
