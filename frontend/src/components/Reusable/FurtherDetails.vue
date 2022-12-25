<template>
  <div id="further_detail">
    <!-- Close Button -->
    <div id="close_btn" @click="closeProjectDetail">X</div>
    <div id="details">
      <!-- Title -->
      <section id="project_title" v-if="detail.title !== ''">
        <h1>{{ detail.title }}</h1>
      </section>
      <!-- Description -->
      <section id="project_desc" v-if="detail.desc !== ''">
        <p v-html="detail.desc"></p>
      </section>
      <!-- Image -->
      <section class="image_container" v-if="detail.images.thumbnail !== ''">
        <img :src="require(`@/assets/imgs/${detail.images.thumbnail}.png`)" />

        <!-- Progress Imgs Including Mobile Designs And / Or Progress Imgs -->
        <div v-if="detail.images.progress.length > 0">
          <div class="mobile_design_container"
            v-for="(progress, progressKey) in detail.images.progress" :key="progressKey">
            <h2>{{ progress.title }}</h2>
            <div class="mobile_design_imgs" ref="mobileDesignVersionRef">
              <img class="mobile_imgs"
                v-for="(progressImg, progressImgKey) in progress.imgs"
                :key="progressImgKey"
                  :src="require(`@/assets/imgs/progress/${progressImg}.png`)" />
            </div>
          </div>
        </div>
      </section>

      <!-- Techs -->
      <section id="tech_container"
        v-if="detail.technologies.length > 0">
        <h2>Technologies</h2>
        <div id="tech_wrapper" ref="techRef">
          <div class="tech_lists"
            v-for="(tech, techKey) in detail.technologies" :key="techKey">
            <a :href="tech.link" target="_blank">
              <h4>{{ tech.name }}</h4>
              <img class="tech_imgs" :src="require(`@/assets/imgs/techs/${tech.img}.png`)" />
            </a>
          </div>
        </div>
      </section>

      <!-- APIs -->
      <section id="api_container"
        v-if="detail.apis.length > 0">
        <h2>APIs</h2>
        <div id="api_wrapper" ref="apiRef">
          <div class="api_lists"
            v-for="(api, apiKey) in detail.apis" :key="apiKey">
            <a v-if="api.name !== 'Internal API' && api.img !== ''"
              :href="api.link" target="_blank">
              <h4>{{ api.name }}</h4>
              <img ref="apiImgRef"
                class="api_imgs" :src="require(`@/assets/imgs/apis/${api.img}.png`)" />
            </a>

            <section v-else>
              <h4>{{ api.name }}</h4>
            </section>
          </div>
        </div>
      </section>

      <section id="links_group" >
        <!-- Videos -->
        <section class="detail" v-if="detail.videos.length !== 0">
          <div class="videos_container" v-for="(video, videoKey) in detail.videos"
            :key="videoKey">
            <h3 class="headers">{{ video.name }}</h3>
            <div class="icon_container">
              <a :href="video.link" target="_blank" alt="demo_video">
                <img src="@/assets/logos/youtube.png" />
              </a>
            </div>
          </div>
        </section>
        <!-- Project Log -->
        <section class="detail" v-if="detail.project_log !== ''">
          <h3 class="headers">Project Log</h3>
          <div class="icon_container">
            <a :href="detail.project_log" target="_blank" alt="report_link">
              <img src="@/assets/logos/project_log.png" />
            </a>
          </div>
        </section>
        <!-- Reports -->
        <section class="detail" v-if="detail.reports.length !== 0">
          <div class="report_container" v-for="(report, reportKey) in detail.reports"
            :key="reportKey">
            <h3 class="headers">{{ report.name }}</h3>
            <div class="icon_container">
              <a :href="report.link" target="_blank" alt="report_link">
                <img src="@/assets/logos/report.png" />
              </a>
            </div>
          </div>
        </section>
        <!-- Github -->
        <section class="detail" v-if="detail.github !== ''">
          <h3 class="headers">Github URL</h3>
          <div class="icon_container">
            <a :href="detail.github" target="_blank" alt="github_url">
              <img src="@/assets/logos/github.png" />
            </a>
          </div>
        </section>
        <!-- Live URL -->
        <section class="detail" v-if="detail.live_url !== ''">
          <h3 class="headers">{{ detail.live_url.name }}</h3>
          <div class="icon_container">
            <a :href="detail.live_url.link" target="_blank" alt="live_url">
              <img src="@/assets/logos/url.png" />
            </a>
          </div>
        </section>
      </section>
    </div>
  </div>
</template>

<script>
import { onMounted, reactive, ref } from 'vue';

export default {
  name: 'FurtherDetail',
  props: {
    emitter: {
      type: Object,
    },
    furtherDetail: {
      type: Object,
    },
    mobile: {
      type: Boolean,
    },
  },
  setup(props) {
    // Props
    const detail = reactive({
      title: props.furtherDetail.title,
      desc: props.furtherDetail.description,
      role: props.furtherDetail.role,
      images: {
        thumbnail: props.furtherDetail.images.thumbnail,
        progress: props.furtherDetail.images.progress,
      },
      videos: props.furtherDetail.extra_links.videos,
      presentations: props.furtherDetail.extra_links.presentations,
      project_log: props.furtherDetail.extra_links.project_log,
      reports: props.furtherDetail.extra_links.reports,
      github: props.furtherDetail.github,
      live_url: props.furtherDetail.live_url,
      technologies: props.furtherDetail.technologies,
      apis: props.furtherDetail.apis,
      organisation: props.furtherDetail.organisation,
    });
    const isMobile = ref(props.mobile);

    // DOM Ref
    const mobileDesignVersionRef = ref(null);
    const techRef = ref(null);
    const apiRef = ref(null);
    const apiImgRef = ref(null);

    function closeProjectDetail() {
      props.emitter.emit('project_card', null);
    }

    function createGridDisplay(target, index, childNode, DOMRef) {
      const newTarget = typeof (target[index][childNode]) !== 'string'
        ? target[index][childNode]
        : target;

      const numOfColsOnMobile = newTarget.length < 6
        ? newTarget.length
        : 3;

      const numOfColsOnBigScreen = newTarget.length < 10
        ? newTarget.length
        : 7;

      // eslint-disable-next-line no-nested-ternary
      const percentage = !isMobile.value
        ? (newTarget.length < 8
          ? `${100 / newTarget.length}%`
          : `${100 / numOfColsOnBigScreen}%`)
        : `${100 / numOfColsOnMobile}%`;

      // eslint-disable-next-line no-unused-expressions
      !isMobile.value
        ? DOMRef.setAttribute('style',
          `grid-template-columns: repeat(${numOfColsOnBigScreen}, ${percentage})`)
        : DOMRef.setAttribute('style',
          `grid-template-columns: repeat(${numOfColsOnMobile}, ${percentage})`);
    }

    onMounted(() => {
      if (detail.images.progress.length > 0) {
        for (let i = 0; i < detail.images.progress.length; i += 1) {
          createGridDisplay(detail.images.progress, i, 'imgs', mobileDesignVersionRef.value[0]);
        }
      }

      if (detail.technologies.length > 0) {
        for (let i = 0; i < detail.technologies.length; i += 1) {
          createGridDisplay(detail.technologies, i, 'img', techRef.value);
        }
      }

      if (detail.apis.length > 0) {
        if (detail.apis.length === 1 && detail.apis[0].img !== '') {
          // eslint-disable-next-line no-unused-expressions
          !isMobile.value
            ? apiImgRef.value[0].setAttribute('style', 'width: 15%')
            : apiImgRef.value[0].setAttribute('style', 'width: 25%');
        }
        for (let i = 0; i < detail.apis.length; i += 1) {
          createGridDisplay(detail.apis, i, 'img', apiRef.value);
        }
      }
    });

    return {
      closeProjectDetail,
      detail,
      mobileDesignVersionRef,
      techRef,
      apiRef,
      apiImgRef,
    };
  },
};
</script>

<style lang="scss" scoped>
#further_detail {
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

  #details {
    position: relative;
    width: 100%;
    margin: 10px;
    padding: 20px;
    color: rgb(0, 0, 0);
    border-radius: 5px;
    background-color: rgba(87, 87, 87, 0.233);
    text-align: left;

    .image_container {
      position: relative;
      width: 100%;
      margin: 0;

      img {
        width: 100%;
        border-radius: 10px;
        box-shadow: 0px 0px 10px black;
      }

      .mobile_design_container {
        margin: 45px 0;

        .mobile_design_imgs {
          display: grid;
          margin-top: 10px;

          .mobile_imgs {
            padding: 10px;
            width: 100%;
          }
        }
      }
    }

    #tech_container {
      position: relative;
      margin-top: 55px;
      padding: 10px;
      border-radius: 5px;
      background-color: rgba(185, 130, 202, 0.25);

      #tech_wrapper,
      #api_wrapper {
        display: grid;

        .tech_lists,
        .api_lists {
          margin-top: 10px;
          width: 65%;
          position: relative;
          left: 50%;
          transform: translateX(-50%);
          cursor: pointer;
          text-decoration: none;
          transition: all .2s;

          a {
            text-decoration: none;
            color: black;
          }

          .tech_imgs {
            padding: 10px;
            width: 50%;
          }

          &:hover {
            text-decoration: underline;
          }
        }
      }
    }

    #api_container {
      @extend #tech_container;
      background-color: rgba(202, 177, 130, 0.25);

      .api_imgs {
        padding: 10px;
        width: 50%;
      }
    }

    section {
      margin: 20px;
      text-align: center;
    }

    #project_title {
      text-align: center;
      background-color: rgba(255, 250, 250, 0.75);
      padding: 5px;
      border-radius: 5px;
    }
    #project_desc {
      text-align: left;
    }
    #links_group {
      position: relative;
      display: flex;
      flex-direction: row;
      align-items: center;
      justify-content: center;

      .detail {
        position: relative;
        width: 50%;

        .headers {
          margin: 5px;
        }
      }
    }

    .icon_container {
      position: relative;
      width: 100px;
      cursor: pointer;
      left: 50%;
      transform: translateX(-50%);

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

/* Responsive Image */
@media screen and (max-width: 770px) {
  #details {
    font-size: 12px;

    section {
      margin: 10px !important;
    }

    .image_container {
      .mobile_design_container {
        margin: 45px 0 !important;

        .mobile_design_imgs {
          .mobile_imgs {
            padding: 10px !important;
          }
        }
      }
    }

    #tech_container,
    #api_container {
      position: relative;
      margin-top: 55px;
      background-color: rgba(201, 202, 130, 0.25);

      #tech_wrapper,
      #api_wrapper {
        .tech_lists,
        .api_lists {
          margin-top: 10px !important;
          width: 60% !important;

          .tech_imgs {
            padding: 10px !important;
            width: 50% !important;
          }
        }
      }
    }

    #links_group {
      margin: 10px !important;

      .icon_container {
        width: 75% !important;
      }
    }
  }
}
@media screen and (max-width: 600px) {
  #details {
    font-size: 10px;

    section {
      margin: 10px !important;
    }

    .image_container {
      .mobile_design_container {
        margin: 20px 0 !important;

        .mobile_design_imgs {
          .mobile_imgs {
            padding: 5px !important;
          }
        }
      }
    }

    #tech_container,
    #api_container {
      position: relative;
      margin-top: 55px;
      background-color: transparent !important;

      #tech_wrapper,
      #api_wrapper {
        .tech_lists,
        .api_lists {
          margin-top: 10px !important;
          width: 60% !important;

          .tech_imgs {
            padding: 10px !important;
            width: 50% !important;
          }
        }
      }
    }

    #links_group {
      margin: 10px !important;

      .icon_container {
        width: 50% !important;
      }
    }
  }
}
@media screen and (max-width: 460px) {
  #details {
    font-size: 7px;

    section {
      margin: 8px !important;
    }

    .image_container {
      .mobile_design_container {
        margin: 20px 0 !important;

        .mobile_design_imgs {
          .mobile_imgs {
            padding: 2px !important;
          }
        }
      }
    }

    #tech_container,
    #api_container {
      position: relative;
      margin-top: 25px !important;
      background-color: transparent !important;

      #tech_wrapper,
      #api_wrapper {
        .tech_lists,
        .api_lists {
          margin-top: 10px !important;
          width: 100% !important;

          .tech_imgs {
            padding: 10px !important;
            width: 75% !important;
          }
        }
      }
    }

    #links_group {
      margin: 2.5px !important;

      .icon_container {
        width: 25% !important;
      }
    }
  }
}
</style>
