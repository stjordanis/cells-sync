import React, {Component, Fragment} from 'react'
import {CompoundButton} from "office-ui-fabric-react";
import SyncTask from "./SyncTask";
import {Route} from 'react-router-dom'
import {Translation} from "react-i18next";

class TasksList extends Component {

    render() {
        const {syncTasks, sendMessage, onDelete} = this.props;
        return (
            <div>
                <Translation>{(t) =>
                    <Route render={({history}) =>
                        <Fragment>
                            <div>
                                {Object.keys(syncTasks).map(k => {
                                    const task = syncTasks[k];
                                    return <SyncTask
                                        key={k}
                                        state={task}
                                        sendMessage={sendMessage.bind(this)}
                                        openEditor={() => {
                                            history.push('/edit/' + task.Config.Uuid)
                                        }}
                                        onDelete={onDelete.bind(this)}
                                    />
                                })}
                            </div>
                            <div style={{padding: 20, textAlign: 'center'}}>
                                <CompoundButton
                                    iconProps={{iconName: 'Add'}}
                                    secondaryText={t('main.create.legend')}
                                    onClick={() => {
                                        history.push('/create')
                                    }}
                                >{t('main.create')}</CompoundButton>
                            </div>
                        </Fragment>
                    }/>
                }</Translation>
            </div>
        );
    }

}

export default TasksList