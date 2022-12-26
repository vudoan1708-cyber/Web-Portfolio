<template>
  <div id="nav">
    <ul id="els">
      <li class="nav_el" :class="selected === 'home' ? 'selected' : 'not_selected'"
          @click="() => { selectNavItem('home') }">
        <a href="#intro">Home</a>
      </li>
      <li class="nav_el" :class="selected === 'portfolio' ? 'selected' : 'not_selected'"
          @click="() => { selectNavItem('portfolio') }">
        <a href="#content">Portfolio</a>
      </li>
      <li class="nav_el" :class="selected === 'blog' ? 'selected' : 'not_selected'"
          @click="() => { selectNavItem('blog') }">
        <a href="#blogs">Blog</a>
      </li>
      <li class="nav_el" :class="selected === 'resume' ? 'selected' : 'not_selected'"
          @click="() => { selectNavItem('resume') }">
        <a href="#resume">Resume</a>
      </li>
    </ul>
  </div>
</template>

<script>
import { ref } from 'vue';

export default {
  name: 'Navigation',
  setup() {
    const selected = ref('home');

    const selectNavItem = (item) => {
      selected.value = item;
    };

    return {
      selected,
      selectNavItem,
    };
  },
};
</script>

<style lang="scss" scoped>
#nav {
  display: flex;
  justify-content: space-around;
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  background-color: black;
  border-bottom-left-radius: 25%;
  border-bottom-right-radius: 25%;
  z-index: 1;

  #els {
    display: grid;
    grid-template-columns: repeat(4, 25%);
    list-style: none;
    padding: 20px;

    .nav_el {
      margin: 20px;
      color: rgb(218, 218, 218);
      font-size: 1.25rem;
      cursor: pointer;

      a {
        text-decoration: none;
        color: rgb(218, 218, 218);
        font-weight: normal;
        transition: color .2s ease-in-out, border-bottom .2s ease-in-out;

        &:hover {
          color: rgb(226, 91, 91);
        }
      }
    }
    .nav_el.selected {
      a {
        padding-bottom: 8px;
        border-bottom: 2px solid #e25b5b;
        color: rgb(226, 91, 91);
      }
    }
  }
}

/* Responsive Text */
@media screen and (max-width: 770px) {
  #nav {
    width: 100%;

    #els {
      padding: 20px;

      .nav_el {
        font-size: 20px !important;
      }
    }
  }
}
@media screen and (max-width: 600px) {
  #nav {
    width: 100%;

    #els {
      padding: 20px;

      .nav_el {
        font-size: 18px !important;
      }
    }
  }
}
@media screen and (max-width: 460px) {
  #nav {
    #els {
      padding: 12px;

      .nav_el {
        font-size: 12px !important;
      }
    }
  }
}
</style>
